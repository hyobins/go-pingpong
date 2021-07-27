package api

import (
	"github.com/hyobins/go-pingpong/service"
	"github.com/labstack/echo"
)

type UserController struct {
	service.UserService
}

func (UserController *UserController) PongHandler(c echo.Context) error {
	result, err := UserController.UserService.PongHandler()
}

