package handlers

import (
	"Bookify/internal/models"
	"Bookify/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	services services.BookService
}

func NewBookHandler(services services.BookService) *BookHandler {
	return &BookHandler{services: services}
}

func (h *BookHandler) GetAllBook(c echo.Context) error {
	books,err := h.services.GetAllBook()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Failed to Fetch"})
	}
	return c.JSON(http.StatusOK,books)

}

func (h *BookHandler) GetBookById(c echo.Context) error {
	id ,err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error":"Invalid Input"})
	}
	book ,err:= h.services.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK,book)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	newBook := new(models.Book);
	if err := c.Bind(&newBook); err != nil {
		c.JSON(http.StatusBadRequest,map[string]string{"error":"Invalid Input"});
	}

	if err := h.services.CreateBook(newBook); err != nil {
		c.JSON(http.StatusInternalServerError,map[string]string{"error":"Failed to create book"})
	}

	return c.JSON(http.StatusCreated,newBook)
}