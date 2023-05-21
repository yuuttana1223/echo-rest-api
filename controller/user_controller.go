package controller

import (
	"echo-rest-api/model"
	"echo-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type UserControllerImpl struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) UserController {
	return &UserControllerImpl{uc}
}

func (c *UserControllerImpl) SignUp(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	resUser, err := c.userUseCase.SignUp(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, resUser)
}

func (c *UserControllerImpl) Login(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	tokenString, err := c.userUseCase.Login(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		// Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	ctx.SetCookie(&cookie)
	return ctx.NoContent(http.StatusOK)
}

func (c *UserControllerImpl) Logout(ctx echo.Context) error {
	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		// Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	ctx.SetCookie(&cookie)
	return ctx.NoContent(http.StatusOK)
}
