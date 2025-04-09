package routes

import (
	"BookStore/internal/delivery"
	"BookStore/internal/repository"
	"BookStore/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	as := service.NewAuthService(ur)
	ah := delivery.NewAuthHandler(as)

	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", ah.Register)
		authRoutes.POST("/login", ah.Login)
	}
}
