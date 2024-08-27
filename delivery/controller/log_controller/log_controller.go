package log_controller

import (
	"loan-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	logUsecase domain.LogUsecase
}

func NewLogController(logUsecase domain.LogUsecase) *LogController {
	return &LogController{
		logUsecase: logUsecase,
	}
}

func (lc *LogController) GetLogs(c *gin.Context) {
	eventType := c.Query("event")
	order := c.DefaultQuery("order", "desc")

	logs, err := lc.logUsecase.GetLogs(c.Request.Context(), eventType, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
