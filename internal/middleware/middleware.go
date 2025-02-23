package middleware

import (
	"Bookify/internal/config"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWT Middleware
func JWTMiddleware(secret string) echo.MiddlewareFunc {
	cfk := config.AppConfigInstance.Key
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfk.SecretKey),
		TokenLookup: "header:Authorization",
	})
}
