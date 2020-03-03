package courses

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
	"github.com/seuc-frp-utn/api/middlewares"
	"github.com/seuc-frp-utn/api/roles"
)

var (
	CourseController *application.IController
	CourseService *application.IService
	courseRepository *application.IRepository
)

func initialize() {
	courseRepository = NewRepository(database.Db)
	CourseService = NewService(courseRepository)
	CourseController = NewController(CourseService)
	err := database.Migrate(database.Db, Course{})
	if err != nil {
		panic(err)
	}
}

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.GET("/:uuid", middlewares.UUID, (*CourseController).Read)
	group.GET("/", (*CourseController).ReadAll)
	group.Use(middlewares.JWT)
	{
		group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*CourseController).Create)
		group.PUT("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*CourseController).Update)
		group.DELETE("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*CourseController).Remove)
	}
	return group
}