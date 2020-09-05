package main

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"

	"net/http"
	"net/http/httputil"
	"os"
)

var graphqlServerUrl = os.Getenv("GRAPHQL_SERVER")

func GraphQL() func(c *gin.Context) {
	base, err := (&url.URL{}).Parse(graphqlServerUrl)
	if err != nil {
		panic(err)
	}
	p := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = base.Scheme
			req.URL.Host = base.Host
			req.URL.Path = strings.Replace(req.URL.Path, "gql/", "", 1)
		},
	}
	return func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	}
}
