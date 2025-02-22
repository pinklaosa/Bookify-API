package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseHandler struct {
	context echo.Context
}

func DefineContext(context echo.Context) *BaseHandler {
	return &BaseHandler{context: context}
}

func (b *BaseHandler) JSONBadRequest(err error) error {
	return b.context.JSON(http.StatusBadRequest,
		map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err.Error()})
}

func (b *BaseHandler) JSONInternalServer(err error) error {
	return b.context.JSON(http.StatusInternalServerError,
		map[string]interface{}{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
}

func (b *BaseHandler) JSONSuccessRequest(data interface{}) error {
	return b.context.JSON(http.StatusOK,
		map[string]interface{}{
			"status": http.StatusOK,
			"data":   data,
		})
}

func (b *BaseHandler) JSONCreated(data interface{}) error {
	return b.context.JSON(http.StatusCreated,
		map[string]interface{}{
			"status": http.StatusOK,
			"data":   data,
		})
}

func (b *BaseHandler) JSONNotFound(err error) error {
	return b.context.JSON(http.StatusNotFound,
		map[string]interface{}{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})
}
