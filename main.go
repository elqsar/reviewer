package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"reviewer/repo"
	"reviewer/web"
	userInfo "reviewer/web/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET was not provided")
	}

	var username string
	var database string
	var poolSize int
	flag.StringVar(&username, "username", "", "Database user username")
	flag.StringVar(&database, "database", "", "Database name")
	flag.IntVar(&poolSize, "poolSize", 5, "Db pool size")
	flag.Parse()

	router := echo.New()

	pool := repo.NewPool(username, database, poolSize)
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
	router.Use(middleware.CORS())

	router.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	router.POST("/users", usersHandler.CreateUser)
	router.POST("/auth", authHandler.Auth)

	api := router.Group("/api")
	api.Use(middleware.JWT([]byte(secret)))
	api.Use(userInfo.UserInfo)
	api.GET("/reviews", reviewHandler.GetAllReviews)
	api.DELETE("/reviews/:id", reviewHandler.Delete)
	api.GET("/users/:id/reviews", reviewHandler.GetUsersReviews)
	api.POST("/users/:id/reviews", reviewHandler.Create)

	log.Fatalf("Server was unable to start: %s", router.Start(":3000"))
}
