package services

import (
	"Bookify/internal/models"
	"Bookify/internal/repositories"
)

type BookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBook(params *models.GetBookParams) ([]models.Book, error) {
	return s.repo.GetAllBook(params)
}

func (s *BookService) GetBookById(id int) (*models.Book, error) {
	return s.repo.GetBookById(id)
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) DeleteBookById(id int) error {
	return s.repo.DeleteBookById(id);
}

func (s *BookService) UpdateBook(id int,book *models.Book) error {
	return s.repo.UpdateBook(id,book);
}