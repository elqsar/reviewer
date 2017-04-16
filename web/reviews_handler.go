package web

import (
	"github.com/labstack/echo"
	"net/http"
	"reviewer/entity"
	"reviewer/repo"
)

type ReviewsHandler struct {
	Reviews *repo.ReviewsRepo
}

func (handler *ReviewsHandler) Create(ctx echo.Context) error {
	review := &entity.Review{}
	if err := ctx.Bind(review); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userId := ctx.Get("userId").(int)
	review.User = userId
	id, err := handler.Reviews.CreateReview(review)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	review.Id = id
	return ctx.JSON(http.StatusCreated, review)
}

func (handler *ReviewsHandler) GetAll(ctx echo.Context) error {
	userId := ctx.Get("userId").(int)
	reviews, err := handler.Reviews.GetReviews(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, reviews)
}
