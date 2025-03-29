package service

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
