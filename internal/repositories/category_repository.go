package repositories

import (
	"Bookify/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) GetCategoryById(id int) (*models.Category, error) {
	var categories models.Category
	result := r.db.First(&categories, id)
	return &categories, result.Error
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	result := r.db.Create(category)
	return result.Error
}

func (r *CategoryRepository) UpdateCategory(id int, updatedCategory *models.Category) error {
	var category models.Category
	if err := r.db.First(&category,id).Error; err != nil {
		return err;
	}
	return r.db.Model(&category).Updates(updatedCategory).Error;
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	var category models.Book
	if err := r.db.First(&category,id).Error; err != nil {
		return err
	}
	return r.db.Delete(&models.Category{},id).Error

}