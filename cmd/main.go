package main

import (
	"Bookify/internal/handlers"
	"Bookify/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	BookHandler := handlers.NewBookHandler()
	routes.RegisterBookAllRoutes(e,BookHandler)
	e.Logger.Fatal(e.Start(":1324"))
}
