package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"reviewer/entity"
	"reviewer/repo"
	"time"
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
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	claims["sub"] = user.Id
	claims["email"] = user.Email

	signed, err := token.SignedString([]byte("supersecret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": signed})
}
