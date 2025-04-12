package service

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   repository.UserRepository
	jwtKey []byte
}

func NewAuthService(repo repository.UserRepository, jwtKey string) *AuthService {
	return &AuthService{repo: repo, jwtKey: []byte(jwtKey)}
}

func (uc *AuthService) Register(username, email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     models.RoleUser,
	}
	err = uc.repo.Create(user)
	if err != nil {
		return "", err
	}

	token := generateJwtToken(user)

	return token.SignedString(uc.jwtKey)
}

func (uc *AuthService) Login(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := generateJwtToken(user)

	return token.SignedString(uc.jwtKey)
}

func generateJwtToken(user *models.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"user_role": user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	return token
}
