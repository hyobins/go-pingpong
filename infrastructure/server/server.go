package server

import (
	"github.com/hyobins/go-pingpong/controller/api"
	"github.com/hyobins/go-pingpong/repository"
	"github.com/hyobins/go-pingpong/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// type Server struct {
// 	MainDB *gorm.DB
// }

func Init() {
	e := echo.New()

	// api controller setting
	userController := InjectUserController()
	userController.Init(e)

	// swagger setting
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// func (server Server) InjectDb() *gorm.DB {
// 	return server.MainDB
// }

//ioc

func InjectUserRepository() *repository.UserRepositoryImpl {
	return repository.UserRepositoryImpl{}.NewUserRepositoryImpl()
}

func InjectUserService() *service.UserServiceImpl {
	return service.UserServiceImpl{}.NewUserServiceImpl(InjectUserRepository())
}

func InjectUserController() *api.UserController {
	return api.UserController{}.NewUserController(InjectUserService())
}
