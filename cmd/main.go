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

	CategoryRepository := repositories.NewCategoryRepository(db)
	CategoryService := services.NewCategoryService(*CategoryRepository);
	CategoryHandler := handlers.NewCategoryHandler(*CategoryService);

	UserRepository := repositories.NewUserRepository(db)
	AuthService := services.NewAuthService(*UserRepository)
	UserHandler := handlers.NewAuthHandler(*AuthService)

	
	e := echo.New()
	routes.RegisterUserAllRoutes(e,UserHandler)
	
	protected := e.Group("");
	routes.RegisterBookAllRoutes(protected,BookHandler)
	routes.RegisterCategoryAllRoutes(e,CategoryHandler);
	
	e.Logger.Fatal(e.Start(":1324"))
}
