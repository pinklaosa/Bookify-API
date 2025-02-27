package services

import (
	"Bookify/internal/models"
	"Bookify/internal/repositories"
)

type ReviewService struct {
	repo repositories.ReviewRepository
}

func (s *ReviewService) GetBookReview(bookID int) (*models.Review,error) {
	return s.repo.GetBookReview(bookID)
}