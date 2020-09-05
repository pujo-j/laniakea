package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

var logger *zap.Logger

type Storage interface {
	Exists(ctx context.Context, path string) (bool, error)
	Get(ctx context.Context, path string) ([]byte, error)
	Put(ctx context.Context, path string, data []byte) error
}

func main() {
	var err error
	if os.Getenv("DEBUG") != "" {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level.SetLevel(zap.DebugLevel)
		logger, _ = cfg.Build()
		logger.Info("Starting in debug mode")
	} else {
		logger, _ = zap.NewProduction()
	}
	var store Storage
	dbUrl := os.Getenv("DB_URL")
	bucketUrl := os.Getenv("BUCKET_URL")
	if bucketUrl != "" {
		store, err = initCloudStorage(bucketUrl)
		if err != nil {
			logger.Panic("connecting to cloud storage", zap.Error(err))
		}
		logger.Info("Using cloud storage :" + bucketUrl)
	} else if dbUrl != "" {
		store, err = InitDb(dbUrl)
		if err != nil {
			logger.Panic("connecting to database", zap.Error(err))
		}
		logger.Info("Using postgresql storage :" + dbUrl)
	}
	var port = "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/*filepath", func(c *gin.Context) {
		fp := c.Param("filepath")
		logger.Debug("GET", zap.String("path", fp))
		bytes, err := store.Get(c.Request.Context(), fp)
		if err != nil {
			c.Writer.WriteHeader(404)
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Write(bytes)
	})
	router.PUT("/*filepath", func(c *gin.Context) {
		fp := c.Param("filepath")
		logger.Debug("PUT", zap.String("path", fp))
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Writer.WriteHeader(400)
			return
		}
		logger.Debug("inserting blob", zap.String("id", fp), zap.Int("size", len(bytes)))
		err = store.Put(c.Request.Context(), fp, bytes)
		if err != nil {
			logger.Error("inserting blob", zap.Error(err))
			c.Writer.WriteHeader(500)
			return
		}
		c.Writer.WriteHeader(201)
		return
	})
	router.HEAD("/*filepath", func(c *gin.Context) {
		fp := c.Param("filepath")
		logger.Debug("HEAD", zap.String("path", fp))
		exists, err := store.Exists(c.Request.Context(), fp)
		if err != nil {
			logger.Error("checking for blob", zap.Error(err))
			c.Writer.WriteHeader(500)
		} else {
			if exists {
				c.Writer.WriteHeader(200)
				return
			} else {
				c.Writer.WriteHeader(404)
				return
			}
		}
	})
	logger.Info("Listening on 0.0.0.0:" + port)
	logger.Error("Listening on 0.0.0.0:"+port, zap.Error(router.Run(":"+port)))
}
