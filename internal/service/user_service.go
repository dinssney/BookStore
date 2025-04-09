package service

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
