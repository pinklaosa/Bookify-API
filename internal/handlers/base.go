package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type BaseHandler struct{}

// NewBaseHandler returns a new instance of BaseHandler
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// JSONBadRequest returns a 400 Bad Request response
func (b *BaseHandler) JSONBadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, Response{
		Status: http.StatusBadRequest,
		Error:  err.Error(),
	})
}

// JSONInternalServer returns a 500 Internal Server Error response
func (b *BaseHandler) JSONInternalServer(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, Response{
		Status: http.StatusInternalServerError,
		Error:  err.Error(),
	})
}

// JSONSuccess returns a 200 OK response
func (b *BaseHandler) JSONSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   data,
	})
}

// JSONCreated returns a 201 Created response
func (b *BaseHandler) JSONCreated(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, Response{
		Status: http.StatusCreated,
		Data:   data,
	})
}

// JSONNotFound returns a 404 Not Found response
func (b *BaseHandler) JSONNotFound(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, Response{
		Status: http.StatusNotFound,
		Error:  err.Error(),
	})
}

func (b *BaseHandler) JSONUnAuth(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, Response{
		Status: http.StatusUnauthorized,
		Error:  err.Error(),
	})
}
