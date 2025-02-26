package routes

import (
	"Bookify/internal/config"
	"Bookify/internal/handlers"
	"Bookify/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterCategoryRoutes(pub *echo.Group, h *handlers.CategoryHandler, cfk config.KeyConfig)  {
	pub.Use(middleware.JWTMiddleware(cfk.SecretKey))
	pub.GET("/categories",h.GetAllCategory)
	pub.GET("/categories/:id",h.GetCategoryById)

}


func RegisterCategoryAdminRoutes(admin *echo.Group, h *handlers.CategoryHandler, cfk config.KeyConfig)  {
	admin.Use(middleware.JWTMiddleware(cfk.SecretKey))
	admin.Use(middleware.RoleMiddleware("admin"))
	admin.POST("/categories",h.CreateCategory)
	admin.PUT("/categories/:id",h.UpdateCategory)
	admin.DELETE("/categories/:id",h.DeleteCategory)
}


func RegisterCategoryAllRoutes(pub, admin *echo.Group, categoryHandler *handlers.CategoryHandler) {
	cfk := config.AppConfigInstance.Key
	RegisterCategoryRoutes(pub, categoryHandler, cfk)
	RegisterCategoryAdminRoutes(admin, categoryHandler, cfk)
}