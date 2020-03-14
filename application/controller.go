package application

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/utils"
	"net/http"
	"reflect"
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

type Controller struct {
	service *IService
}

func NewController(service *IService) *IController {
	var c IController
	c = &Controller{
		service: service,
	}
	return &c
}


func (c Controller) GetService() (*IService, error) {
	if c.service == nil {
		return nil, errors.New("undefined service")
	}
	return c.service, nil
}

func (c *Controller) SetService(service *IService) error {
	c.service = service
	return nil
}

func (c *Controller) Create(typeOf interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		t := reflect.TypeOf(typeOf)

		if t.Kind() != reflect.Struct {
			ctx.AbortWithError(http.StatusInternalServerError, errors.New("input body is not a struct"))
			return
		}

		var body map[string]interface{}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		entity := reflect.Indirect(reflect.New(t))
		if !entity.CanAddr() {
			ctx.AbortWithError(http.StatusInternalServerError, errors.New("not addressable"))
		}

		entity = utils.FillStruct(entity, body)

		result, err := (*c.service).Create(entity)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusCreated, result)
	}
}

func (c *Controller) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	result, err := (*c.service).Get(uuid)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Controller) GetAll(ctx *gin.Context) {
	result, err := (*c.service).GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *Controller) Update(typeOf interface{}) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")

		v := reflect.New(reflect.TypeOf(typeOf))

		if err := ctx.ShouldBindJSON(v.Interface()); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		result, err := (*c.service).Update(uuid, v.Elem())

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func (c *Controller) Remove(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	result, err := (*c.service).Get(uuid)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
