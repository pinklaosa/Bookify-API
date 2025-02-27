package handlers

import (
	"Bookify/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	base BaseHandler
	services services.ReviewService
}

func NewReviewHandler(services *services.ReviewService) *ReviewHandler{
	return &ReviewHandler{
		base: *NewBaseHandler(),
		services: *services}
}

func (h *ReviewHandler) GetBookReview(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("book_id"));
	if err != nil {
		return h.base.JSONBadRequest(c,err)
	}
	reviews, err := h.services.GetBookReview(bookId);
	if err != nil {
		return h.base.JSONNotFound(c,err);
	}
	return h.base.JSONSuccess(c,reviews)
}

