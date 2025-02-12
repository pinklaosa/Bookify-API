package handlers

import (
	"Bookify/internal/services"
	"net/http"

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