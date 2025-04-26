package routes

import (
	"BookStore/internal/delivery"
	"BookStore/internal/middleware"
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"BookStore/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserBooksRoutes(router *gin.Engine, db *gorm.DB, jwtKey string) {
	r := repository.NewUserBooksRepository(db)
	s := service.NewUserBooksService(r)
	h := delivery.NewUserBooksHandler(s)

	userBooks := router.Group("/api/v1/user-books")
	userBooks.Use(middleware.AuthMiddleware(jwtKey))
	userBooks.Use(middleware.RequireRole(models.RoleUser))
	{
		userBooks.GET("", h.GetUserBooks)
		userBooks.POST("/:bookId", h.AddBookToUser)
		userBooks.DELETE("/:bookId", h.RemoveBookFromUser)
	}
}
