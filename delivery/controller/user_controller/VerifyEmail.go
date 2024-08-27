package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) VerifyEmail(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}
	uc.userUsecase.VerifyEmail(context.Background(), tokenString)
	c.JSON(http.StatusCreated, gin.H{"message": "Email has been successfully verified"})
}
