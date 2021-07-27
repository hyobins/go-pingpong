package server

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/hyobins/go-pingpong/repository"
)

type Server struct {
	MainDB *gorm.DB
}

func (server Server) Init(){
	e := echo.New()

	// api controller setting
	userController := server.InjectUserController()
	userController.Init

	e.Logger.Fatal(e.Start(":8080"))
}

func (server Server) InjectDb() *gorm.DB {
	return server.MainDB
}

func (server Server) InjectUserRepository() *repository.UserRepositoryImpl {
	return repository.UserRepositoryImpl{}.NewUserRepositoryImpl(server.InjectDb())
}

func (server Server) InjectUserService() *service.UserServiceImpl {
	return service.UserServiceImpl{}.NewUserServiceImpl(server.InjectUserRepository())
}

func (server Server) InjectUserController() *api.UserController {
	return api.UserController{}.NewUserController(server.InjectUserService())
}
