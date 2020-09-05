package main

import (
	"context"
	"github.com/dgraph-io/ristretto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"sync"
	"time"
)

type DbStorage struct {
	pool  *pgxpool.Pool
	cache *ristretto.Cache
}

var updates = UpdateBatch{
	updates: make(map[string]time.Time),
	lock:    sync.Mutex{},
}

type UpdateBatch struct {
	updates map[string]time.Time
	lock    sync.Mutex
}

func (d *DbStorage) AsyncMarkAccess() {
	update := func() {
		updates.lock.Lock()
		defer updates.lock.Unlock()
		if len(updates.updates) > 0 {
			batch := &pgx.Batch{}
			toUpdate := make([]string, 0)
			for path, update := range updates.updates {
				batch.Queue("UPDATE public.STORAGE SET LAST_ACCESS=$1 WHERE ID=$2", update, path)
				toUpdate = append(toUpdate, path)
			}
			results := d.pool.SendBatch(context.Background(), batch)
			for _, path := range toUpdate {
				_, err := results.Exec()
				if err != nil {
					logger.Error("updating lastaccess", zap.Error(err))
				} else {
					delete(updates.updates, path)
				}
			}
		}
	}
	for {
		time.Sleep(10 * time.Second)
		update()
	}
}

func InitDb(dbUrl string) (Storage, error) {
	pool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	rows, err := pool.Query(context.Background(), "SELECT 1 FROM STORAGE")
	if err != nil {
		_, err = pool.Exec(context.Background(), `
			CREATE TABLE public.STORAGE
			(
				ID character varying COLLATE pg_catalog."default" NOT NULL,
				DATA bytea NOT NULL,
				CREATED timestamp with time zone,
				LAST_ACCESS timestamp with time zone,
				CONSTRAINT STORAGE_PK PRIMARY KEY (ID)
			)`)
		if err != nil {
			return nil, err
		}
	} else {
		rows.Close()
	}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 100000,
		MaxCost:     128 * 1024 * 1024,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	res := &DbStorage{
		pool:  pool,
		cache: cache,
	}
	go res.AsyncMarkAccess()
	return res, nil
}

func (d *DbStorage) Exists(ctx context.Context, path string) (bool, error) {
	_, ok := d.cache.Get(path)
	if ok {
		return true, nil
	} else {
		row := d.pool.QueryRow(ctx, "SELECT ID FROM STORAGE WHERE ID=$1", path)
		var id string
		err := row.Scan(&id)
		if err != nil {
			if err == pgx.ErrNoRows {
				return false, nil
			} else {
				return false, err
			}
		}
		return true, nil
	}
}

func (d *DbStorage) Get(ctx context.Context, path string) ([]byte, error) {
	cached, ok := d.cache.Get(path)
	if ok {
		data := cached.([]byte)
		d.cache.Set(path, data, int64(len(data)))
		updates.lock.Lock()
		defer updates.lock.Unlock()
		updates.updates[path] = time.Now()
		return data, nil
	} else {
		row := d.pool.QueryRow(ctx, "SELECT DATA FROM STORAGE WHERE ID=$1", path)
		var data []byte
		err := row.Scan(&data)
		if err != nil {
			return nil, err
		}
		d.cache.Set(data, data, int64(len(data)))
		updates.lock.Lock()
		defer updates.lock.Unlock()
		updates.updates[path] = time.Now()
		return data, nil
	}
}

func (d *DbStorage) Put(ctx context.Context, path string, data []byte) error {
	_, err := d.pool.Exec(ctx, `
		INSERT INTO STORAGE(ID,DATA, CREATED,LAST_ACCESS) VALUES($1,$2,$3,$3)
`, path, data, time.Now())
	if err != nil {
		return err
	}
	d.cache.Set(path, data, int64(len(data)))
	return nil
}
