package services

import (
	"Bookify/internal/models"
	"Bookify/internal/repositories"
)

type ReviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) GetBookReview(bookId int) (*models.Review,error) {
	return s.repo.GetBookReview(bookId)
}

func (s *ReviewService) CreateReview(review *models.Review) error{
	return s.repo.CreateReview(review)
}
