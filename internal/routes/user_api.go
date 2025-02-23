package routes

import (
	"Bookify/internal/handlers"

	"github.com/labstack/echo/v4"
)


func RegisterUserRoutes(e *echo.Echo,h *handlers.AuthHandler){
	e.POST("/register",h.Register);
	e.POST("/login",h.Login)
}

func RegisterUserAllRoutes(e *echo.Echo,userHandler *handlers.AuthHandler) {
	RegisterUserRoutes(e,userHandler)
}