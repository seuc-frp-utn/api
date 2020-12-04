package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/pkg/application"
	"github.com/seuc-frp-utn/api/pkg/auth"
	"github.com/seuc-frp-utn/api/pkg/database"
	"github.com/seuc-frp-utn/api/pkg/middlewares"
	"github.com/seuc-frp-utn/api/pkg/roles"
	"time"
)

var (
	UserController *application.IController
	UserService *application.IService
	userRepository *application.IRepository
)

func initialize() {
	userRepository = NewRepository(database.Db)
	UserService = NewService(userRepository)
	UserController = application.NewController(UserService)
	err := database.Migrate(database.Db, User{})
	if err != nil {
		panic(err)
	}
}

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	group.Use(middlewares.JWT)
	{
		group.GET("/:uuid", middlewares.UUID, middlewares.Roles(roles.USER|roles.OPERATOR|roles.ADMIN), (*UserController).Get)
		group.GET("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).GetAll)
		group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Create(UserCreate{}))
		group.PUT("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Update(UserUpdate{}))
		group.DELETE("/:uuid", middlewares.UUID, middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Remove)
	}
	return group
}

func dropTable() {
	database.Drop(User{})
}

func addTestData() {
	uuid := auth.GenerateUUID()
	password, _ := auth.GeneratePassword("root")
	admin := User{
		UUID:  uuid,
		FirstName:  "Admin",
		MiddleName: nil,
		LastName:   "Root",
		Email:      "root@admin.com",
		Birthday:   time.Time{},
		Password:   password,
		Role:       roles.ADMIN|roles.USER|roles.TEACHER,
		DNI:		22476549,
	}

	userUUID := auth.GenerateUUID()
	userPassword, _ := auth.GeneratePassword("12345")
	user := User{
		UUID:  userUUID,
		FirstName:  "User",
		MiddleName: nil,
		LastName:   "Sudo",
		Email:      "user@sudo.com",
		Birthday:   time.Time{},
		Password:   userPassword,
		Role:       roles.USER,
		DNI:		35457889,
	}

	if userRepository != nil {
		db, err := (*userRepository).GetDatabase()
		if err != nil {
			panic(err)
			return
		}
		db.Model(&User{}).Save(&admin)
		db.Model(&User{}).Save(&user)
	}
}

func RegisterDirectTest(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	dropTable()
	addTestData()
	group.GET("/:uuid", (*UserController).Get)
	group.GET("/", (*UserController).GetAll)
	group.POST("/", (*UserController).Create(UserCreate{}))
	group.PUT("/:uuid", (*UserController).Update(UserUpdate{}))
	group.DELETE("/:uuid", (*UserController).Remove)
	return group

}

func RegisterTestJWT(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	dropTable()
	addTestData()
	group.Use(middlewares.JWT)
	{
		group.GET("/:uuid", (*UserController).Get)
		group.GET("/", (*UserController).GetAll)
		group.POST("/", (*UserController).Create(UserCreate{}))
		group.PUT("/:uuid", (*UserController).Update(UserUpdate{}))
		group.DELETE("/:uuid", (*UserController).Remove)
	}
	return group
}

func RegisterTestRoles(group *gin.RouterGroup) *gin.RouterGroup {
	initialize()
	dropTable()
	addTestData()
	group.GET("/:uuid", middlewares.Roles(roles.USER|roles.OPERATOR|roles.ADMIN), (*UserController).Get)
	group.GET("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).GetAll)
	group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Create(UserCreate{}))
	group.PUT("/:uuid", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Update(UserUpdate{}))
	group.DELETE("/:uuid", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Remove)
	return group

}