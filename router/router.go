package router

import (
	"echo-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(c controller.UserController) *echo.Echo {
	e := echo.New()
	e.POST("/sign-up", c.SignUp)
	e.POST("/login", c.Login)
	e.POST("/logout", c.Logout)
	return e
}
