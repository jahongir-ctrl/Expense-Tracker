package controller

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateRecurringExpenseHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	var input struct {
		Amount      int    `json:"amount" binding:"required"`
		Category    string `json:"category" binding:"required"`
		Description string `json:"description" binding:"required"`
		Frequency   string `json:"frequency" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input body", "details": err.Error()})
		return
	}

	exp := models.Recurringexpense{
		UserID:      userID,
		Category:    input.Category,
		Description: input.Description,
		Frequency:   input.Frequency,
		Amount:      input.Amount,
		NextDate:    time.Now().AddDate(0, 0, 1),
		Active:      true,
	}
	if err := service.CreateRecurringExpense(exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recurring expense", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Recurring expense created"})
}

func GetRecurringExpensesHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	data, err := repository.GetRecurringExpenses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recurring expense", "details": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": "Recurring expenses", "data": data})
	return
}

func DeleteRecurringExpenseHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteRecurringExpense(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Recurring expense deleted"})
}

func StopRecurringExpenseHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := service.StopRecurringExpense(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop recurring expense", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Recurring expense stopped"})
}
