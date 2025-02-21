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

func (s *CategoryService) GetCategoryById(id int) (*models.Category,error){
	return s.repo.GetCategoryById(id);
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.repo.CreateCategory(category);
}

func (s *CategoryService) UpdateCategory(id int,updatedCategory *models.Category) error  {
	return s.repo.UpdateCategory(id,updatedCategory);
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id);
}