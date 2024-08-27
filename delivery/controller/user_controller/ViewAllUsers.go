package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) ViewAllUsers(c *gin.Context) {
	users, err := uc.userUsecase.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
