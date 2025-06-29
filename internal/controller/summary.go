package controller

import (
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBalanceSummaryHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	summary, err := service.GetBalanceSummary(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance summary", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
