package service

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"errors"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	CreateBook(book *models.Book, publisherID uint) error
	UpdateBook(book *models.Book, publisherID uint, id uint) error
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

func (s *bookService) CreateBook(book *models.Book, publisherID uint) error {
	book.PublisherID = publisherID
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(req *models.Book, publisherID uint, id uint) error {
	book, err := s.GetBookByID(id)
	if err != nil {
		return err
	} else if book.PublisherID != publisherID {
		return errors.New("you are not allowed to update this book")
	}

	book.Title = req.Title
	book.Author = req.Author
	book.Description = req.Description
	book.CoverImageUrl = req.CoverImageUrl

	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
