package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/courses"
	"github.com/seuc-frp-utn/api/diplomas"
	"github.com/seuc-frp-utn/api/login"
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
		// Users routes
		//////////////////////////////////////////////////
		usersGroup := v1.Group("/users")
		{
			users.Register(usersGroup)
		}
		//////////////////////////////////////////////////
		// Courses routes
		//////////////////////////////////////////////////
		coursesGroup := v1.Group("/courses")
		{
			courses.Register(coursesGroup)
		}
		//////////////////////////////////////////////////
		// Diplomas routes
		//////////////////////////////////////////////////
		diplomasGroup := v1.Group("/diplomas")
		{
			diplomas.Register(diplomasGroup)
		}
		//////////////////////////////////////////////////
		// Payments routes
		//////////////////////////////////////////////////
		payments := v1.Group("/payments")
		{
			payments.GET("/")
		}
		//////////////////////////////////////////////////
		// Authentication routes
		//////////////////////////////////////////////////
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register")
			authGroup.POST("/login", login.Handler)
		}
	}
	return v1
}