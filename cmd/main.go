package main

import (
	"BookStore/internal/db"
	"BookStore/internal/delivery"
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"BookStore/internal/routes"
	"BookStore/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	db := db.Instance

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := delivery.NewBookHandler(bookService)

	db.AutoMigrate(models.Book{})
	
	r := gin.Default()
	routes.RegisterBookRoutes(r, bookHandler)

	r.Run(":8080")
}
