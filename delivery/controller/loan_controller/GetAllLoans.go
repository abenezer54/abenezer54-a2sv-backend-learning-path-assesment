package loan_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) GetAllLoans(c *gin.Context) {
	// Extract query parameters
	status := c.DefaultQuery("status", "all")
	order := c.DefaultQuery("order", "")

	// Call the use case
	loans, err := lc.loanUsecase.GetAllLoans(c.Request.Context(), status, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}
