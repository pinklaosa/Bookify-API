package repositories

import (
	"Bookify/internal/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAllBook() ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *BookRepository) GetBookById(id int) (*models.Book , error) {
	var books models.Book
	result := r.db.First(&books,id)
	return &books , result.Error
}
