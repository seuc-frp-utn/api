package main

import "github.com/gin-gonic/gin"

func registerRoutes(group *gin.RouterGroup) *gin.RouterGroup {
	v1 := group.Group("/1.0")
	{
		users := v1.Group("/users")
		{
			users.GET("/")
		}

		courses := v1.Group("/courses")
		{
			courses.GET("/")
		}

		certificates := v1.Group("/certificates")
		{
			certificates.GET("/")
		}

		payments := v1.Group("/payments")
		{
			payments.GET("/")
		}
	}
	return v1
}