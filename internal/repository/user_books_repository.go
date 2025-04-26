package repository

import (
	"BookStore/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserBooksRepository struct {
	db *gorm.DB
}

func NewUserBooksRepository(db *gorm.DB) *UserBooksRepository {
	return &UserBooksRepository{db: db}
}

// AddBookToUser adds a book to user's book collection
func (r *UserBooksRepository) AddBookToUser(userID, bookID uint) error {
	// Check if user exists
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Check if book exists
	var book models.Book
	if err := r.db.First(&book, bookID).Error; err != nil {
		return errors.New("book not found")
	}

	// Check if the association already exists
	var count int64
	r.db.Table("user_books").
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Count(&count)

	if count > 0 {
		return errors.New("book already added to user's collection")
	}

	// Add the association
	return r.db.Exec("INSERT INTO user_books (user_id, book_id) VALUES (?, ?)", userID, bookID).Error
}

// RemoveBookFromUser removes a book from user's book collection
func (r *UserBooksRepository) RemoveBookFromUser(userID, bookID uint) error {
	// Check if user exists
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Check if book exists
	var book models.Book
	if err := r.db.First(&book, bookID).Error; err != nil {
		return errors.New("book not found")
	}

	// Check if the association exists
	var count int64
	r.db.Table("user_books").
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Count(&count)

	if count == 0 {
		return errors.New("book not found in user's collection")
	}

	// Remove the association
	return r.db.Exec("DELETE FROM user_books WHERE user_id = ? AND book_id = ?", userID, bookID).Error
}

func (r *UserBooksRepository) GetUserBooks(userID uint) ([]models.Book, error) {
	// Check if user exists
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var books []models.Book
	err := r.db.Joins("JOIN user_books ON user_books.book_id = books.id").
		Where("user_books.user_id = ?", userID).
		Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}
