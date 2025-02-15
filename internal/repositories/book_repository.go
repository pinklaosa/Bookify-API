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

func (r *BookRepository) GetBookById(id int) (*models.Book, error) {
	var books models.Book
	result := r.db.First(&books, id)
	return &books, result.Error
}

func (r *BookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) DeleteBookById(id int) error {
	var book models.Book;
	if err := r.db.First(&book,id).Error; err != nil {
		return err;
	}
	return r.db.Delete(&models.Book{},id).Error
}

func (r *BookRepository) UpdateBook(id int,updatedBook *models.Book) error {
	var book models.Book;
	if err := r.db.First(&book,id).Error; err != nil {
		return err;
	}
	return r.db.Model(&book).Updates(updatedBook).Error;

}