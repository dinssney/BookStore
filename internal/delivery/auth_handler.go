package delivery

import (
	"BookStore/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	token, err := ah.authService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(201, gin.H{"token": token})
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	token, err := ah.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
