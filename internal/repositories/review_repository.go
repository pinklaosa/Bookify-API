package repositories

import (
	"Bookify/internal/models"
	"errors"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetBookReview(bookId int) (*models.Review, error) {
	var reviews models.Review
	if err := r.db.Where("book_id = ?",bookId).Find(&reviews).Error; err != nil {
		return nil , errors.New("no review")
	}
	return &reviews,nil
}

func (r *ReviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) DeleteReview(id int) error{
	var review models.Review
	if err := r.db.First(&review,id).Error; err != nil {
		return err
	}
	return r.db.Delete(&models.Review{},id).Error
}