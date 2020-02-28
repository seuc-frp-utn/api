package application

import "github.com/gin-gonic/gin"

type IController interface {
	GetService() (*IService, error)
	SetService(service *IService) error
	Create(ctx *gin.Context)
	Read(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
}