package api

import (
	"net/http"

	"github.com/hyobins/go-pingpong/service"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service.UserService
}

func (UserController) NewUserController(svc service.UserService) *UserController {
	return &UserController{svc}
}

func (userController *UserController) Init(e *echo.Echo) {
	e.GET("/ping", userController.PongHandler)
	e.GET("", exampleHandler)
}

// PongHandler return "pong!"
// @Summary Pong
// @Description ddd
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /users [get]
func (userController *UserController) PongHandler(c echo.Context) error {
	result, err := userController.UserService.PongHandler()
	if err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Info(result)
	return c.String(http.StatusOK, result)
}

func exampleHandler(c echo.Context) error {
	return c.String(http.StatusOK, "default routing")
}
