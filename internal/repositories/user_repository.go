package repositories

import (
	"Bookify/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

//Register
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

