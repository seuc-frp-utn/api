package users

import (
	"github.com/seuc-frp-utn/api/application"
	"github.com/seuc-frp-utn/api/database"
)

var (
	UserController *application.IController
	UserService *application.IService
	UserRepository *application.IRepository
)

func init() {
	UserRepository = NewRepository(database.Db)
	UserService = NewService(UserRepository)
	UserController = NewController(UserService)
}
