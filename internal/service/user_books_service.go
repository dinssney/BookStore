package service

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
)

type UserBooksService struct {
	repo *repository.UserBooksRepository
}

func NewUserBooksService(repo *repository.UserBooksRepository) *UserBooksService {
	return &UserBooksService{repo: repo}
}

func (s *UserBooksService) AddBookToUser(userID, bookID uint) error {
	return s.repo.AddBookToUser(userID, bookID)
}

func (s *UserBooksService) GetUserBooks(userID uint) ([]models.Book, error) {
	return s.repo.GetUserBooks(userID)
}

func (s *UserBooksService) RemoveBookFromUser(userID, bookID uint) error {
	return s.repo.RemoveBookFromUser(userID, bookID)
}
