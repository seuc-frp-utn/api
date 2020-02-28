package main

import "github.com/gin-gonic/gin"

func initializeRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes := router.Group("/")
	{
		registerRoutes(routes)
	}
	return router
}
