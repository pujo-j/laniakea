package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var storageServer = os.Getenv("STORAGE_SERVER")

func Storage() func(c *gin.Context) {
	base, err := (&url.URL{}).Parse(storageServer)
	if err != nil {
		panic(err)
	}
	p := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = base.Scheme
			req.URL.Host = base.Host
			req.URL.Path = strings.Replace(req.URL.Path, "storage/", "", 1)
		},
	}
	return func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	}
}
