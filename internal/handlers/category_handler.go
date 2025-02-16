package handlers

import (
	"Bookify/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	services services.CategoryService
}

func NewCategoryHandler(services services.CategoryService) *CategoryHandler {
	return &CategoryHandler{services: services}
}

func (h *CategoryHandler) GetAllCategory(c echo.Context) error  {
	result, err := h.services.GetAllCategory();
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Failed tpo Fetch"})
	}
	return c.JSON(http.StatusOK, result);
}