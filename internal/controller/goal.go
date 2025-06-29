package controller

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateGoalHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	var input struct {
		Title        string  `json:"title" binding:"required"`
		TargetAmount float64 `json:"target_amount" binding:"required"`
		Description  string  `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	goal := models.Goal{
		UserID:       userID,
		Title:        input.Title,
		TargetAmount: input.TargetAmount,
		Description:  input.Description,
	}

	if err := service.CreateGoal(goal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create goal", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Goal created successfully", "goal": goal})
}

func UpdateGoalHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	goalID, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		TargetAmount float64 `json:"target_amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.UpdateGoalAmount(goalID, userID, input.TargetAmount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update goal amount", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Goal amount updated successfully"})
}

func DeleteGoalHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	goalID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteGoal(goalID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete goal", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Goal deleted successfully"})
}

func GetGoalsHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	goals, err := service.GetGoalsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve goals", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"goals": goals})
}
