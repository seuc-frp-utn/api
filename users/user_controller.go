package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"net/http"
)

type Controller struct {
	service *application.IService
}

func NewController(service *application.IService) *application.IController {
	var c application.IController
	c = &Controller{
		service: service,
	}
	return &c
}

func (c Controller) GetService() (*application.IService, error) {
	if c.service == nil {
		return nil, errors.New("undefined service")
	}
	return c.service, nil
}

func (c *Controller) SetService(service *application.IService) error {
	c.service = service
	return nil
}

func (c Controller) Create(ctx *gin.Context) {
	var body UserCreate

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := (*c.service).Create(body)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}

func (c Controller) Read(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	result, err := (*c.service).Get(uuid)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Controller) ReadAll(ctx *gin.Context) {
	result, err := (*c.service).GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c Controller) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	var body UserUpdate
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := (*c.service).Update(uuid, body)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)


}

func (c Controller) Remove(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	result, err := (*c.service).Get(uuid)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}