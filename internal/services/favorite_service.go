package services

import (
	"Bookify/internal/models"
	"Bookify/internal/repositories"
)

type FavoriteService struct {
	repo *repositories.FavoriteRepository
}

func NewFavoriteService(repo *repositories.FavoriteRepository) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (s *FavoriteService) AddFavorite(userID, bookID int) error {
	favorite := &models.Favorite{
		UserID: userID,
		BookID: bookID,
	}
	return s.repo.AddFavorite(favorite)
}

func (s *FavoriteService) RemoveFavorite(userID, bookID int) error {
	return s.repo.RemoveFavorite(userID, bookID)
}

func (s *FavoriteService) GetUserFavorites(userID int) ([]models.Favorite, error) {
	return s.repo.GetUserFavorites(userID)
}
