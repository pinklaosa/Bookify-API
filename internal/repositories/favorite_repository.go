package repositories

import (
	"Bookify/internal/models"

	"gorm.io/gorm"
)

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
    return &FavoriteRepository{db: db}
}

func (r *FavoriteRepository) AddFavorite(favorite *models.Favorite) error {
    return r.db.Create(favorite).Error
}

func (r *FavoriteRepository) RemoveFavorite(userID, bookID int) error {
    return r.db.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.Favorite{}).Error
}

func (r *FavoriteRepository) GetUserFavorites(userID int) ([]models.Favorite, error) {
    var favorites []models.Favorite
    err := r.db.Where("user_id = ?", userID).Find(&favorites).Error
    return favorites, err
}