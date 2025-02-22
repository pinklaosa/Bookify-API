package handlers

import (
	"Bookify/internal/models"
	"Bookify/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	base BaseHandler
	services services.CategoryService
}

func NewCategoryHandler(services services.CategoryService) *CategoryHandler {
	return &CategoryHandler{ 
		base: *NewBaseHandler(),	
		services: services}
}

func (h *CategoryHandler) GetAllCategory(c echo.Context) error {
	result, err := h.services.GetAllCategory()
	if err != nil {
		return h.base.JSONBadRequest(c, err)
	}
	return h.base.JSONSuccess(c, result)
}

func (h *CategoryHandler) GetCategoryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.base.JSONBadRequest(c, err)
	}
	category, err := h.services.GetCategoryById(id)
	if err != nil {
		return h.base.JSONNotFound(c, err)
	}
	return h.base.JSONSuccess(c, category)
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(&category); err != nil {
		return h.base.JSONBadRequest(c, err)
	}

	if err := h.services.CreateCategory(category); err != nil {
		return h.base.JSONInternalServer(c, err)
	}

	return h.base.JSONCreated(c, category)
}

func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.base.JSONBadRequest(c, err)
	}

	updatedCategory := new(models.Category)
	if err := c.Bind(&updatedCategory); err != nil {
		return h.base.JSONBadRequest(c, err)
	}
	if err := h.services.UpdateCategory(id, updatedCategory); err != nil {
		return h.base.JSONNotFound(c, err)
	}
	return h.base.JSONSuccess(c, updatedCategory)
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.base.JSONBadRequest(c, err)
	}

	if err := h.services.DeleteCategory(id); err != nil {
		return h.base.JSONInternalServer(c, err)
	}

	return h.base.JSONSuccess(c, "Deleted")
}
