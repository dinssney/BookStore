package middleware

import (
	"BookStore/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRole middleware checks if the user has the required role
func RequireRole(requiredRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user role from the context (set by AuthMiddleware)
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
