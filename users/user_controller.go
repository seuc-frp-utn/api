package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
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

}

func (c Controller) Read(ctx *gin.Context) {

}

func (c Controller) Update(ctx *gin.Context) {

}

func (c Controller) Remove(ctx *gin.Context) {

}