package controller

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateIncomeHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	var input struct {
		Amount      float64 `json:"amount" binding:"required"`
		Source      string  `json:"source" binding:"required"`
		Description string  `json:"description" binding:"required"`
		//Date        string  `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//date, _ := time.Parse(time.RFC3339, input.Date)

	income := models.Income{
		UserID:      userID,
		Amount:      input.Amount,
		Source:      input.Source,
		Description: input.Description,
		//Date:        date.Format(time.RFC3339),
	}

	err := service.CreateIncome(userID, input.Amount, input.Source, input.Description, time.Now().Format(time.RFC3339))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create income", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Income created successfully", "income": income})
}

func GetIncomesHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	incomes, err := service.GetIncomes(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve incomes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"incomes": incomes})
}

func DeleteIncomeHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	incomeID, err := strconv.Atoi(c.Param("id"))

	err = service.DeleteIncome(incomeID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete income", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Income deleted successfully"})
}
