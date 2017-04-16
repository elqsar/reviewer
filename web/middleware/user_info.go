package middleware

import (
	"github.com/labstack/echo"
	"reviewer/repo"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func UserInfo(repo *repo.UsersRepo) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			username := claims["username"].(string)
			result, err := repo.GetByUsername(username)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}
			ctx.Set("userId", result.Id)
			return next(ctx)
		}
	}
}
