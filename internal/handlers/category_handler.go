package handlers

import (
	"Bookify/internal/models"
	"Bookify/internal/services"
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Failed to Fetch"})
	}
	return c.JSON(http.StatusOK, result);
}

func (h *CategoryHandler) GetCategoryById(c echo.Context) error  {
	id,err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"});
	}

	category, err := h.services.GetCategoryById(id);
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error":"Category Not Found"});
	}

	return c.JSON(http.StatusOK,category)
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
	}

	if err := h.services.CreateCategory(category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Failed to create category"})
	}

	return c.JSON(http.StatusCreated, category);
}

func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	updatedCategory := new(models.Category);

	if err := c.Bind(&updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
	}

	if err := h.services.UpdateCategory(id,updatedCategory); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Category updated successfully"})
}	

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"));
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error":"Invalid Input"});
	}

	if err := h.services.DeleteCategory(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Fail to Delete"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success":"Deleted"})
}
