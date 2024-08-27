package loan_controller

import (
	"loan-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) UpdateLoanStatus(c *gin.Context) {
	var req domain.UpdateLoanStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	updatedLoan, err := lc.loanUsecase.UpdateLoanStatus(c.Request.Context(), id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLoan)
}
