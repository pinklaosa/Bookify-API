package services

import (
	"Bookify/internal/models"
	"Bookify/internal/repositories"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo};
}

func (s *CategoryService) GetAllCategory() ([]models.Category,error) {
	return s.repo.GetAllCategory();
}