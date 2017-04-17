package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"reviewer/repo"
	"reviewer/web"
	userInfo "reviewer/web/middleware"
)

func main() {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET was not provided")
	}
	var username string
	var database string
	flag.StringVar(&username, "username", "", "Database user username")
	flag.StringVar(&database, "database", "", "Database name")
	flag.Parse()

	router := echo.New()

	pool := repo.NewReviewerPool(username, database, 5)
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
	protected.Use(middleware.JWT([]byte(secret)))
	protected.Use(userInfo.UserInfo)
	protected.GET("/reviews", reviewHandler.GetAllReviews)
	protected.GET("/users/:id/reviews", reviewHandler.GetUsersReviews)
	protected.POST("/users/:id/reviews", reviewHandler.Create)

	router.Start(":3000")
}
