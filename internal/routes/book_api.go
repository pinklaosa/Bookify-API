package routes

import (
	"Bookify/internal/config"
	"Bookify/internal/handlers"
	"Bookify/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterBookRoutes(e *echo.Group, h *handlers.BookHandler) {
	cfk := config.AppConfigInstance.Key
	e.Use(middleware.JWTMiddleware(cfk.SecretKey))
	e.GET("/books", h.GetAllBook)
	e.GET("/books/:id", h.GetBookById)
	e.POST("/books", h.CreateBook)
	e.DELETE("/books/:id", h.DeleteBookById)
	e.PUT("/books/:id", h.UpdateBook)
}

func RegisterBookAllRoutes(e *echo.Group, bookHandler *handlers.BookHandler) {

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	RegisterBookRoutes(e, bookHandler)
}
