package handlers

import (
	"Bookify/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FavoriteHandler struct {
	base BaseHandler
	service *services.FavoriteService
}

func NewFavoriteHandler(service *services.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		base: *NewBaseHandler(),
		service: service}
}


func (h *FavoriteHandler) AddFavorite(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	userID := c.Get("user_id").(int)
	if err := h.service.AddFavorite(userID, bookID); err != nil {
		return h.base.JSONBadRequest(c,err)
	}
	return h.base.JSONCreated(c,"Added")
}


func (h *FavoriteHandler) RemoveFavorite(c echo.Context) error {
	userID := c.Get("user_id").(int)
	bookID, _ := strconv.Atoi(c.Param("book_id"))

	if err := h.service.RemoveFavorite(userID, bookID); err != nil {
		return h.base.JSONInternalServer(c,err)
	}
	return h.base.JSONSuccess(c,"Favorite removed")
}


func (h *FavoriteHandler) GetUserFavorites(c echo.Context) error {
	userID := c.Get("user_id").(int)
	favorites, err := h.service.GetUserFavorites(userID)
	if err != nil {
		return h.base.JSONInternalServer(c,err)
	}

	return h.base.JSONSuccess(c,favorites)
}