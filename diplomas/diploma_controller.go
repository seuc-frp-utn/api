package diplomas

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"net/http"
)

type IController interface {
	application.IController
	Find(ctx *gin.Context)
}

type Controller struct {
	application.IController
}

func NewController(service *application.IService) *IController {
	var c IController
	base := application.NewController(service)

	c = &Controller{
		IController: *base,
	}
	return &c
}

func (c Controller) Find(ctx *gin.Context) {
	token := ctx.Param("uuid")
	var service *application.IService
	var err error

	if service, err = c.GetService(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	var result interface{}
	result, err = (*service).Find("token", token)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
