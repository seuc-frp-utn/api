package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initializeRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(Cors())

	routes := router.Group("/")
	{
		registerRoutes(routes)
	}
	return router
}

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)
}