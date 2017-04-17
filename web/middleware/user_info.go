package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func UserInfo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		stringId, ok := claims["sub"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		userId, err := strconv.Atoi(stringId)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		ctx.Set("userId", userId)
		return next(ctx)
	}
}
