package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/auth"
	"github.com/seuc-frp-utn/api/database"
	"github.com/seuc-frp-utn/api/middlewares"
	"github.com/seuc-frp-utn/api/roles"
	"time"
)

var (
	UserController *application.IController
	UserService *application.IService
	userRepository *application.IRepository
)

func initialize(test bool) {
	userRepository = NewRepository(database.Db)
	UserService = NewService(userRepository)
	UserController = NewController(UserService)
	err := database.Migrate(database.Db, User{})
	if err != nil {
		panic(err)
	}
	if test {
		dropTable()
		addTestData()
	}
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
		Role:       roles.ADMIN|roles.USER,
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

func Register(group *gin.RouterGroup) *gin.RouterGroup {
	initialize(true)
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
	initialize(true)
	group.GET("/:uuid", (*UserController).Read)
	group.GET("/", (*UserController).ReadAll)
	group.POST("/", (*UserController).Create)
	group.PUT("/:uuid", (*UserController).Update)
	group.DELETE("/:uuid", (*UserController).Remove)
	return group

}

func RegisterTestJWT(group *gin.RouterGroup) *gin.RouterGroup {
	initialize(true)
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
	initialize(true)
	group.GET("/:uuid", middlewares.Roles(roles.USER|roles.OPERATOR|roles.ADMIN), (*UserController).Read)
	group.GET("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).ReadAll)
	group.POST("/", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Create)
	group.PUT("/:uuid", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Update)
	group.DELETE("/:uuid", middlewares.Roles(roles.OPERATOR|roles.ADMIN), (*UserController).Remove)
	return group

}