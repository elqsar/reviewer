package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"reviewer/repo"
	"reviewer/web"
)

func main() {
	router := echo.New()

	users := repo.NewUsersRepo()
	usersHandler := &web.UsersHandler{
		Users: users,
	}
	authHandler := &web.AuthHandler{
		Users: users,
	}

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	router.POST("/users", usersHandler.CreateUser)
	router.POST("/auth", authHandler.Auth)

	protected := router.Group("/api")
	protected.Use(middleware.JWT([]byte("supersecret")))
	protected.GET("/restricted", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{"result": "success"})
	})

	router.Start(":3000")
}
