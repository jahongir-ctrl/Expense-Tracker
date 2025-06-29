package controller

import (
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDailyReportHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	report, err := service.GetDailyReport(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch daily report", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}

func GetWeeklyReportHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	report, err := service.GetWeeklyReport(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weekly report", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}

func GetMonthlyReportHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	report, err := service.GetMonthlyReport(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monthly report", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}
