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
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	result, err := (*c.service).Create(body)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}

func (c Controller) Read(ctx *gin.Context) {
	panic("implement me")
}

func (c Controller) ReadAll(ctx *gin.Context) {
	panic("implement me")
}

func (c Controller) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c Controller) Remove(ctx *gin.Context) {
	panic("implement me")
}