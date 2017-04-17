package web

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"reviewer/core"
	"reviewer/entity"
	"reviewer/repo"
)

type AuthHandler struct {
	Users *repo.UsersRepo
}

func (handler *AuthHandler) Auth(ctx echo.Context) error {
	user := &entity.User{}
	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	result, err := handler.Users.GetByUsername(user.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	signed, err := core.CreateToken(result)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": signed})
}
