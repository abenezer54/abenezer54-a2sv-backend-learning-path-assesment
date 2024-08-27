package loan_controller

import (
	"loan-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) ApplyForLoan(c *gin.Context) {
	var req domain.LoanApplicationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := lc.loanUsecase.ApplyForLoan(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
