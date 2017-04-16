package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"reviewer/repo"
	"reviewer/web"
	userInfo "reviewer/web/middleware"
	"flag"
)

func main() {
	var username string
	var database string
	flag.StringVar(&username, "username", "", "Database user username")
	flag.StringVar(&database, "database", "", "Database name")
	flag.Parse()

	router := echo.New()

	pool := repo.NewReviewerPool(username, database)
	users := repo.NewUsersRepo(pool)
	reviews := repo.NewReviewsRepo(pool)

	usersHandler := &web.UsersHandler{
		Users: users,
	}
	authHandler := &web.AuthHandler{
		Users: users,
	}
	reviewHandler := &web.ReviewsHandler{
		Reviews: reviews,
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
	protected.Use(userInfo.UserInfo(users))
	protected.GET("/reviews", reviewHandler.GetAll)
	protected.POST("/reviews", reviewHandler.Create)

	router.Start(":3000")
}
