package services

import (
	"Bookify/internal/config"
	"Bookify/internal/models"
	"Bookify/internal/repositories"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Register
func (s *AuthService) Register(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	//check user
	//passed
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	//check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	//genarate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	cfk := config.AppConfigInstance.Key
	secretKey := cfk.SecretKey
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString, nil

}
