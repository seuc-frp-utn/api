package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
	"github.com/seuc-frp-utn/api/middlewares"
	"github.com/seuc-frp-utn/api/roles"
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
	group.Use(middlewares.Handler)
	{
		group.GET("/:uuid", middlewares.Handler(roles.USER|roles.OPERATOR|roles.ADMIN), (*UserController).Read)
		group.GET("/", middlewares.Handler(roles.OPERATOR|roles.ADMIN), (*UserController).ReadAll)
		group.POST("/", middlewares.Handler(roles.OPERATOR|roles.ADMIN), (*UserController).Create)
		group.PUT("/:uuid", middlewares.Handler(roles.OPERATOR|roles.ADMIN), (*UserController).Update)
		group.DELETE("/:uuid", middlewares.Handler(roles.OPERATOR|roles.ADMIN), (*UserController).Remove)
	}
	return group
}
