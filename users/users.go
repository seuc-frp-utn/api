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
	database.DBMigrate(database.Db, User{})
}

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.Use(middlewares.JWT)
	{
		group.GET("/:uuid", middlewares.UUID, middlewares.Roles(roles.USER|roles.OPERATOR|roles.ADMIN), (*UserController).Read)
		group.GET("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).ReadAll)
		group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Create)
		group.PUT("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Update)
		group.DELETE("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Remove)
	}
	return group
}

func RegisterDirectTest(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.GET("/:uuid", (*UserController).Read)
	group.GET("/", (*UserController).ReadAll)
	group.POST("/", (*UserController).Create)
	group.PUT("/:uuid", (*UserController).Update)
	group.DELETE("/:uuid", (*UserController).Remove)
	return group

}

func RegisterTestJWT(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.Use(middlewares.JWT)
	{
		group.GET("/:uuid", (*UserController).Read)
		group.GET("/", (*UserController).ReadAll)
		group.POST("/", (*UserController).Create)
		group.PUT("/:uuid", (*UserController).Update)
		group.DELETE("/:uuid", (*UserController).Remove)
	}
	return group
}

func RegisterTestRoles(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.GET("/:uuid", middlewares.Roles(roles.USER|roles.OPERATOR|roles.ADMIN), (*UserController).Read)
	group.GET("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).ReadAll)
	group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Create)
	group.PUT("/:uuid", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Update)
	group.DELETE("/:uuid", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Remove)
	return group

}