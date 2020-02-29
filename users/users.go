package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
)

var (
	UserController *application.IController
	UserService *application.IService
	userRepository *application.IRepository
)

func initialize() {
	userRepository = NewRepository(database.Db)
	UserService = NewService(userRepository)
	UserController = NewController(UserService)
}

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.GET("/:uuid", (*UserController).Read)
	group.GET("/", (*UserController).ReadAll)
	group.POST("/", (*UserController).Create)
	group.PUT("/:uuid", (*UserController).Update)
	group.DELETE("/:uuid", (*UserController).Remove)
	return group
}
