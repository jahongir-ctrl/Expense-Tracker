package controller

import (
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetBudgetHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	var input struct {
		Category string  `json:"category" binding:"required"`
		Limit    float64 `json:"limit" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.SetBudget(userID, input.Category, input.Limit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set budget", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Budget set successfully"})
}

func GetBudgetsHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	budgets, err := service.GetBudgets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve budgets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"budgets": budgets})
}

func CheckBudgetHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	category := c.Param("category")

	isExceeded, total, limit, err := service.IsBudgetExceeded(userID, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check budget", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"category": category,
		"total":    total,
		"limit":    limit,
		"exceeded": isExceeded,
	})
}
