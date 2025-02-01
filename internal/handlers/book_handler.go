package handlers

import (
	"Bookify/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var books = []models.Book{
	{ID: 1, Title: "React101", Author: "Tide"},
	{ID: 2, Title: "Go lang101", Author: "Tide"},
	{ID: 3, Title: "Python101", Author: "Tide"},
}

type BookHandler struct {
	books []models.Book
}

func NewBookHandler() *BookHandler {
	return &BookHandler{
		books: books,
	}
}

// CRUD
func (b *BookHandler) GetAllBook(c echo.Context) error {
	return c.JSON(http.StatusOK, b.books)
}

func (b *BookHandler) GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for _, book := range b.books {
		if book.ID == id {
			return c.JSON(http.StatusOK, book)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid input"})
}

func (b *BookHandler) CreateBook(c echo.Context) error {
	newBook := models.Book{}
	if err := c.Bind(&newBook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	newBook.ID = len(b.books) + 1
	b.books = append(b.books, newBook)
	return c.JSON(http.StatusCreated, newBook)
}

func (b *BookHandler) UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}

	for i, book := range books {
		if book.ID == id {
			if err := c.Bind(&books[i]); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
			}
			return c.JSON(http.StatusOK, books[i])
		}

	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})

}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"msg": "Book deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})

}
