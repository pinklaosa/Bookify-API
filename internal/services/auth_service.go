package services

import (
	"Bookify/internal/models"
	"Bookify/internal/repositories"
)

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
    return &AuthService{repo: repo}
}


//Register
func (s *AuthService) Register(user *models.User) error {
	return s.repo.CreateUser(user)
}