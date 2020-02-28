package main

import (
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	routes := router.Group("/")
	{
		registerRoutes(routes)
	}

	router.Run(":8080")
}
