package routes

import (
	"Bookify/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterCategoryRoutes(e *echo.Echo, h *handlers.CategoryHandler)  {
	e.GET("/categories",h.GetAllCategory)
	e.GET("/categories/:id",h.GetCategoryById)
	e.POST("/categories",h.CreateCategory)
	e.PUT("/categories/:id",h.UpdateCategory)
	e.DELETE("/categories/:id",h.DeleteCategory)
}



func RegisterCategoryAllRoutes(e *echo.Echo, categoryHandler *handlers.CategoryHandler) {
	RegisterCategoryRoutes(e,categoryHandler)
}