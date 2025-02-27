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

func (r *ReviewRepository) GetBookReview(bookID int) (*models.Review, error) {
	var reviews models.Review
	if err := r.db.Where("book_id = ?",bookID).Find(&reviews).Error; err != nil {
		return nil , errors.New("no review")
	}
	return &reviews,nil
}
