package application

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	GetService() (*IService, error)
	SetService(service *IService) error
	Create(typeOf interface{}) gin.HandlerFunc
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(typeOf interface{}) gin.HandlerFunc
	Remove(ctx *gin.Context)
}

type IControllerCreate interface {
	Create(typeOf interface{}) gin.HandlerFunc
}

type IControllerGet interface {
	Get(ctx *gin.Context)
}

type IControllerGetAll interface {
	GetAll(ctx *gin.Context)
}

type IControllerUpdate interface {
	Update(typeOf interface{}) gin.HandlerFunc
}

type IControllerRemove interface {
	Remove(ctx *gin.Context)
}