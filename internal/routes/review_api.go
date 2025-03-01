package routes

import (
	"Bookify/internal/config"
	"Bookify/internal/handlers"
	"Bookify/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterReviewRoutes(pub *echo.Group,h *handlers.ReviewHandler, cfk config.KeyConfig)  {
	pub.Use(middleware.JWTMiddleware(cfk.SecretKey))
	pub.GET("/review/:book_id",h.GetBookReview)
}

func RegisterReviewAllRoutes(pub, admin *echo.Group,reviewHandler *handlers.ReviewHandler) {
	cfk := config.AppConfigInstance.Key
	RegisterReviewRoutes(pub, reviewHandler,cfk)
}