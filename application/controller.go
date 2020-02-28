package application

import "github.com/gin-gonic/gin"

type IController interface {
	Create(ctx *gin.Context) (interface{}, error)
	Read(ctx *gin.Context) (interface{}, error)
	Update(ctx *gin.Context) (interface{}, error)
	Remove(ctx *gin.Context) (interface{}, error)
}