package diplomas

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
	"github.com/seuc-frp-utn/api/middlewares"
	"github.com/seuc-frp-utn/api/roles"
)

var (
	DiplomaController *IController
	DiplomaService *application.IService
	diplomaRepository *application.IRepository
)

func initialize() {
	diplomaRepository = NewRepository(database.Db)
	DiplomaService = NewService(diplomaRepository)
	DiplomaController = NewController(DiplomaService)
	err := database.Migrate(database.Db, Diploma{})
	if err != nil {
		panic(err)
	}
}

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.GET("/:uuid", middlewares.UUID, (*DiplomaController).Get)
	group.GET("/", (*DiplomaController).GetAll)
	group.GET("/:uuid/token", (*DiplomaController).Find)
	group.Use(middlewares.JWT)
	{
		group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*DiplomaController).Create(DiplomaCreate{}))
		group.PUT("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*DiplomaController).Update(DiplomaUpdate{}))
		group.DELETE("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*DiplomaController).Remove)
	}
	return group
}