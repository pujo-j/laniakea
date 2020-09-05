package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgraph-io/ristretto"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime"
	"strconv"
	"strings"
)

var fileCache *ristretto.Cache
var indexPage fileCacheEntry

type fileCacheEntry struct {
	mime    string
	buf     []byte
	gzBuf   []byte
	brotBuf []byte
	eTag    string
	exist   bool
	forever bool
}

func accept(c *gin.Context) string {
	// Test for gzip and br, brotly has priority
	accepts := strings.Split(c.Request.Header.Get("accept-encoding"), ",")
	res := ""
	for _, a := range accepts {
		a2 := strings.Trim(a, " ")
		if a2 == "gzip" {
			res = a
		}
		if a2 == "br" {
			return "br"
		}
	}
	return res
}

func (f fileCacheEntry) Serve(c *gin.Context) error {
	inm := c.Request.Header.Get("If-None-Match")
	if inm != "" {
		if inm == f.eTag {
			c.Writer.WriteHeader(304)
			return nil
		}
	}
	c.Writer.Header().Set("Content-Type", f.mime)
	c.Writer.Header().Set("ETag", f.eTag)
	accepted := accept(c)
	if f.forever {
		c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	}
	switch accepted {
	case "br":
		if f.brotBuf != nil {
			c.Writer.Header().Set("Content-Encoding", "br")
			c.Writer.Header().Set("Content-Length", strconv.Itoa(len(f.brotBuf)))
			c.Writer.WriteHeader(200)
			_, err := c.Writer.Write(f.brotBuf)
			return err
		}
	case "gzip":
		if f.gzBuf != nil {
			c.Writer.Header().Set("Content-Encoding", "gzip")
			c.Writer.Header().Set("Content-Length", strconv.Itoa(len(f.gzBuf)))
			c.Writer.WriteHeader(200)
			_, err := c.Writer.Write(f.gzBuf)
			return err
		}
	default:

	}
	c.Writer.WriteHeader(200)
	_, err := c.Writer.Write(f.buf)
	return err
}

func newCacheEntry(file string) fileCacheEntry {
	buf, err := ioutil.ReadFile("./dist" + file)
	if err != nil {
		return fileCacheEntry{exist: false}
	}
	res := fileCacheEntry{
		buf:     buf,
		mime:    "application/octet-stream",
		exist:   true,
		forever: true,
	}
	s := strings.Split(file, ".")
	if len(s) > 1 {
		ext := s[len(s)-1]
		m := mime.TypeByExtension("." + ext)
		if m != "" {
			res.mime = m
		}
	}
	// Find if there are gzip and brotly versions
	gzbuf, err := ioutil.ReadFile("./dist" + file + ".gz")
	if err == nil {
		res.gzBuf = gzbuf
	}
	brbuf, err := ioutil.ReadFile("./dist" + file + ".br")
	if err == nil {
		res.brotBuf = brbuf
	}
	hash := sha256.New()
	hash.Write(buf)
	etag := hex.EncodeToString(hash.Sum(nil))
	res.eTag = etag
	return res
}

func init() {
	var err error
	fileCache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 10000,
		MaxCost:     32 * 1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	indexPage = newCacheEntry("/index.html")
	indexPage.forever = false
	fileCache.Set("/index.html", indexPage, int64(len(indexPage.buf)))
}

func ServeFiles(c *gin.Context) {
	fp := c.Request.URL.Path
	//fp := c.Param("filepath")
	if fp == "/" || fp == "/index.html" {
		err := indexPage.Serve(c)
		if err != nil {
			panic(err)
		}
		return
	}
	entryO, found := fileCache.Get(fp)
	if found {
		entry, ok := entryO.(fileCacheEntry)
		if ok {
			if entry.exist {
				err := entry.Serve(c)
				if err != nil {
					panic(err)
				}
				return
			} else {
				err := indexPage.Serve(c)
				if err != nil {
					panic(err)
				}
				return
			}
		}
	} else {
		entry := newCacheEntry(fp)
		fileCache.Set(fp, entry, int64(len(entry.buf)))
		if entry.exist {
			err := entry.Serve(c)
			if err != nil {
				panic(err)
			}
			return
		} else {
			err := indexPage.Serve(c)
			if err != nil {
				panic(err)
			}
			return
		}
	}
}
