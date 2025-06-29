package controller

import (
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetBudgetHandler godoc
// @Summary      Set budget
// @Description  Set a budget limit for a category
// @Tags         budgets
// @Accept       json
// @Produce      json
// @Param        budget  body  object{category=string,limit=number}  true  "Budget info"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/budgets [post]
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

// GetBudgetsHandler godoc
// @Summary      Get budgets
// @Description  Get all budgets for the user
// @Tags         budgets
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/budgets [get]
func GetBudgetsHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	budgets, err := service.GetBudgets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve budgets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"budgets": budgets})
}

// CheckBudgetHandler godoc
// @Summary      Check budget
// @Description  Check if budget is exceeded for a category
// @Tags         budgets
// @Produce      json
// @Param        category   path      string  true  "Category"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/budgets/status [get]
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
