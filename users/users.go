package users

import (
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
)

var (
	UserController *application.IController
	UserService *application.IService
	userRepository *application.IRepository
)

func init() {
	userRepository = NewRepository(database.Db)
	UserService = NewService(userRepository)
	UserController = NewController(UserService)
}
