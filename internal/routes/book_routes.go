package routes

import (
	"BookStore/internal/delivery"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine, bookHandler *delivery.BookHandler) {
	r.GET("/books", bookHandler.GetAllBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.POST("/books", bookHandler.CreateBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
}
