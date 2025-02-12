package main

import (
	"Bookify/internal/database"
	"Bookify/internal/handlers"
	"Bookify/internal/repositories"
	"Bookify/internal/routes"
	"Bookify/internal/services"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.ConnectDb()
	BookRepository := repositories.NewBookRepository(db)
	BookService := services.NewBookService(*BookRepository)
	BookHandler := handlers.NewBookHandler(*BookService)

	e := echo.New()
	
	routes.RegisterBookAllRoutes(e,BookHandler)
	
	e.Logger.Fatal(e.Start(":1324"))
}
