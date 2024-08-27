package user_controller

import (
	"loan-api/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Register request: ", req)
	resp, err := uc.userUsecase.Register(c.Request.Context(), req, uc.Env.VefifyTokenSecret, uc.Env.VerifyTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
