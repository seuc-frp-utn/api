package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/users"
)

func registerRoutes(group *gin.RouterGroup) *gin.RouterGroup {
	v1 := group.Group("/1.0")
	{
		//////////////////////////////////////////////////
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		//////////////////////////////////////////////////
		//////////////////////////////////////////////////
		groupUsers := v1.Group("/users")
		{
			groupUsers.GET("/", (*users.UserController).Read)
			groupUsers.POST("/", (*users.UserController).Create)
			groupUsers.PUT("/:uuid", (*users.UserController).Update)
			groupUsers.DELETE("/:uuid", (*users.UserController).Remove)
		}
		//////////////////////////////////////////////////
		//////////////////////////////////////////////////
		courses := v1.Group("/courses")
		{
			courses.GET("/")
		}
		//////////////////////////////////////////////////
		//////////////////////////////////////////////////
		certificates := v1.Group("/diploma")
		{
			certificates.GET("/")
		}
		//////////////////////////////////////////////////
		//////////////////////////////////////////////////
		payments := v1.Group("/payments")
		{
			payments.GET("/")
		}
		//////////////////////////////////////////////////
		//////////////////////////////////////////////////
		auth := v1.Group("/auth")
		{
			auth.POST("/register")
			auth.POST("/login")
		}
	}
	return v1
}