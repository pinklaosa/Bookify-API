package repositories

import (
	"Bookify/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db:db}
}

func (r *CategoryRepository)GetAllCategory() ([]models.Category,error) {
	var categories []models.Category
	result := r.db.Find(&categories);
	return categories , result.Error
}