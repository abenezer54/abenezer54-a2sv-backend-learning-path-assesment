package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user role from the context (assumes you store it in the request context after authentication)
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the user has an admin role
		if strings.ToLower(userRole.(string)) != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
			c.Abort()
			return
		}

		// Continue to the next handler if the user is an admin
		c.Next()
	}
}
