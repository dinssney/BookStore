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

func SetupBooksRoutes(r *gin.Engine, db *gorm.DB, jwtKey string) {

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := delivery.NewBookHandler(bookService)

	r.GET("/api/v1/books", bookHandler.GetAllBooks)
	r.GET("/api/v1/books/:id", bookHandler.GetBookByID)

	pr := r.Group("/api/v1/books", middleware.AuthMiddleware(jwtKey), middleware.RequireRole(models.RolePublisher))
	{
		pr.POST("/", bookHandler.CreateBook)
		pr.PUT("/:id", bookHandler.UpdateBook)
		pr.DELETE("/:id", bookHandler.DeleteBook)
	}
}
