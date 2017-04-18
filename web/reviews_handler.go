package web

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"reviewer/entity"
	"reviewer/repo"
	"strconv"
)

const (
	IdParam = "id"
)

type ReviewsHandler struct {
	Reviews *repo.ReviewsRepo
}

type ReviewResponse struct {
	Result interface{} `json:"result"`
}

func (handler *ReviewsHandler) Create(ctx echo.Context) error {
	review := &entity.Review{}
	if err := ctx.Bind(review); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userId, err := strconv.Atoi(ctx.Param(IdParam))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	review.User = userId
	id, err := handler.Reviews.CreateReview(review)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	review.Id = id
	return ctx.JSON(http.StatusCreated, &ReviewResponse{review})
}

func (handler *ReviewsHandler) Delete(ctx echo.Context) error {
	reviewId, err := strconv.Atoi(ctx.Param(IdParam))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := handler.Reviews.DeleteReview(reviewId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, &ReviewResponse{})
}

func (handler *ReviewsHandler) GetUsersReviews(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param(IdParam))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	ctxId := ctx.Get("userId").(int)
	if ctxId != userId {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	reviews, err := handler.Reviews.GetReviews(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &ReviewResponse{reviews})
}

func (handler *ReviewsHandler) GetAllReviews(ctx echo.Context) error {
	reviews, err := handler.Reviews.GetAllReviews()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &ReviewResponse{reviews})
}
