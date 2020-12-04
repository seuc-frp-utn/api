package payments

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/pkg/application"
	"github.com/seuc-frp-utn/api/pkg/database"
	"github.com/seuc-frp-utn/api/pkg/middlewares"
	"github.com/seuc-frp-utn/api/pkg/roles"
)

var (
	PaymentController *application.IController
	PaymentService    *application.IService
	paymentRepository *application.IRepository
)

func initialize() {
	paymentRepository = NewRepository(database.Db)
	PaymentService = NewService(paymentRepository)
	PaymentController = application.NewController(PaymentService)
	err := database.Migrate(database.Db, Payment{})
	if err != nil {
		panic(err)
	}
}

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.Use(middlewares.JWT)
	{
		group.GET("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*PaymentController).Get)
		group.GET("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*PaymentController).GetAll)
		group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*PaymentController).Create(PaymentCreate{}))
		group.PUT("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*PaymentController).Update(PaymentUpdate{}))
		group.DELETE("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*PaymentController).Remove)
	}
	return group
}