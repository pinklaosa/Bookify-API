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

func (s *BookService) GetAllBook() ([]models.Book,error) {
	return s.repo.GetAllBook()
}

func (s *BookService) GetBookById(id int) (*models.Book,error) {
	return s.repo.GetBookById(id)
}