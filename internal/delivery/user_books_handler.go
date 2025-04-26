package delivery

import (
	"BookStore/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserBooksHandler struct {
	service *service.UserBooksService
}

func NewUserBooksHandler(service *service.UserBooksService) *UserBooksHandler {
	return &UserBooksHandler{service: service}
}

func (h *UserBooksHandler) GetUserBooks(c *gin.Context) {
	// Get user ID from context (set by AuthMiddleware)
	userID, _ := c.Get("user_id")
	books, err := h.service.GetUserBooks(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *UserBooksHandler) AddBookToUser(c *gin.Context) {
	// Get user ID from context (set by AuthMiddleware)
	userID, _ := c.Get("user_id")

	// Get book ID from URL parameter
	bookID, err := strconv.ParseUint(c.Param("bookId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = h.service.AddBookToUser(userID.(uint), uint(bookID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}

func (h *UserBooksHandler) RemoveBookFromUser(c *gin.Context) {
	// Get user ID from context (set by AuthMiddleware)
	userID, _ := c.Get("user_id")

	// Get book ID from URL parameter
	bookID, err := strconv.ParseUint(c.Param("bookId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = h.service.RemoveBookFromUser(userID.(uint), uint(bookID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
}
