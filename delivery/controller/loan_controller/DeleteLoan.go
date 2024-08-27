package loan_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) DeleteLoan(c *gin.Context) {
	id := c.Param("id")

	err := lc.loanUsecase.DeleteLoan(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "loan application not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan application deleted successfully"})
}
