package users

import (
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
)

var (
	userController *application.IController
	userService *application.IService
	userRepository *application.IRepository
)

func init() {
	userRepository = NewRepository(database.Db)
	userService = NewService(userRepository)
	userController = NewController(userService)
}
