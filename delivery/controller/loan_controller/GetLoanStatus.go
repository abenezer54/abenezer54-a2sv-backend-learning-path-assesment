package loan_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) GetLoanStatus(c *gin.Context) {
	loanID := c.Param("id")

	loanStatus, err := lc.loanUsecase.GetLoanStatus(c.Request.Context(), loanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loanStatus)
}
