package routes

import (
	"Bookify/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterBookRoutes(e *echo.Echo, h *handlers.BookHandler) {
	e.GET("/books", h.GetAllBook)
	e.GET("/books/:id", h.GetBookById)
	e.POST("/books", h.CreateBook)
	e.DELETE("/books/:id", h.DeleteBookById)
}

func RegisterBookAllRoutes(e *echo.Echo, bookHandler *handlers.BookHandler) {

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	RegisterBookRoutes(e, bookHandler)
}
