package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func main() {
	initLogging()
	initAuth()
	var port = "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	proxy := GraphQL()
	id := Id()
	router.NoRoute(func(c *gin.Context) {
		ServeFiles(c)
	})
	router.GET("/id.js", id)
	router.POST("/_auth", Login)
	gql := router.Group("/gql")
	gql.Use(AuthMiddleware)
	gql.POST("*p1", proxy)
	storageProxy := Storage()
	storage := router.Group("/storage")
	storage.Use(AuthMiddleware)
	storage.Any("*p2", storageProxy)
	logger.Info("Listening on 0.0.0.0:" + port)
	logger.Error("Listening on 0.0.0.0:"+port, zap.Error(router.Run(":"+port)))
}
