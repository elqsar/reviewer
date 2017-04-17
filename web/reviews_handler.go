package web

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"reviewer/entity"
	"reviewer/repo"
	"strconv"
)

type ReviewsHandler struct {
	Reviews *repo.ReviewsRepo
}

func (handler *ReviewsHandler) Create(ctx echo.Context) error {
	review := &entity.Review{}
	if err := ctx.Bind(review); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	review.User = userId
	id, err := handler.Reviews.CreateReview(review)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	review.Id = id
	return ctx.JSON(http.StatusCreated, review)
}

func (handler *ReviewsHandler) GetUsersReviews(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	ctxId := ctx.Get("userId").(int)
	if ctxId != userId {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	reviews, err := handler.Reviews.GetReviews(userId)
	fmt.Println("Reveiws: ", reviews)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, reviews)
}

func (handler *ReviewsHandler) GetAllReviews(ctx echo.Context) error {
	reviews, err := handler.Reviews.GetAllReviews()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, reviews)
}
