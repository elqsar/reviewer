package web

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"reviewer/entity"
	"reviewer/repo"
)

type UsersHandler struct {
	Users *repo.UsersRepo
}

func (handler *UsersHandler) CreateUser(ctx echo.Context) error {
	user := &entity.User{}
	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to parse json")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	user.Password = string(hash)
	if err := handler.Users.Create(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"result": "success"})
}
