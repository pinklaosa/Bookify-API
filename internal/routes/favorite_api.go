package routes

import (
	"Bookify/internal/config"
	"Bookify/internal/handlers"
	"Bookify/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterFavoriteRoutes(pub *echo.Group,h *handlers.FavoriteHandler,cfk config.KeyConfig)  {
	pub.Use(middleware.JWTMiddleware(cfk.SecretKey))
	pub.GET("/favorite/:user_id",h.GetUserFavorites)
	pub.POST("/favorite",h.AddFavorite)
	pub.DELETE("/favorite/:book_id",h.RemoveFavorite)
}

func RegisterFavoriteAllroutes(pub, admin *echo.Group,favoriteHandler *handlers.FavoriteHandler) {
	cfk := config.AppConfigInstance.Key
	RegisterFavoriteRoutes(pub,favoriteHandler, cfk)
}