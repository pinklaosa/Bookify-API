package routes

import (
	"Bookify/internal/config"
	"Bookify/internal/handlers"
	"Bookify/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterBookRoutes(pub *echo.Group, h *handlers.BookHandler, cfk config.KeyConfig) {
	pub.Use(middleware.JWTMiddleware(cfk.SecretKey))
	pub.GET("/books", h.GetAllBook)
	pub.GET("/books/:id", h.GetBookById)

}

func RegisterBookAdminRoutes(admin *echo.Group, h *handlers.BookHandler, cfk config.KeyConfig) {
	admin.Use(middleware.JWTMiddleware(cfk.SecretKey))
	admin.Use(middleware.RoleMiddleware("admin"))
	admin.POST("/books", h.CreateBook)
	admin.DELETE("/books/:id", h.DeleteBookById)
	admin.PUT("/books/:id", h.UpdateBook)
}

func RegisterBookAllRoutes(pub, admin *echo.Group, bookHandler *handlers.BookHandler) {
	cfk := config.AppConfigInstance.Key
	RegisterBookRoutes(pub, bookHandler, cfk)
	RegisterBookAdminRoutes(admin, bookHandler, cfk)
}
