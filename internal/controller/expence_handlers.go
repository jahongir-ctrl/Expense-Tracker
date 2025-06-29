package controller

import (
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateExpenseHandler godoc
// @Summary      Create expense
// @Description  Create a new expense for the user
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        expense  body  object{amount=float64,category=string,description=string}  true  "Expense info"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/expenses [post]
func CreateExpenseHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	var input struct {
		Amount      float64 `json:"amount" binding:"required"`
		Category    string  `json:"category" binding:"required"`
		Description string  `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//expense, err := service.CreateExpense(userID, input.Amount, input.Category, input.Description)
	if err := service.CreateExpense(userID, input.Amount, input.Description, input.Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Expense created successfully"})
}

// GetExpencesHandler godoc
// @Summary      Get all expenses
// @Description  Get all expenses for the user
// @Tags         expenses
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/expenses [get]
func GetExpencesHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	expenses, err := service.GetUserExpenses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve expenses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

// UpdateExpenseHandler godoc
// @Summary      Update expense
// @Description  Update an existing expense
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        id      path    int     true  "Expense ID"
// @Param        expense body    object{amount=float64,category=string,description=string}  true  "Expense info"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/expenses/{id} [put]
func UpdateExpenseHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	idStr := c.Param("id")
	expenseID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	var input struct {
		Amount      float64 `json:"amount" binding:"required"`
		Category    string  `json:"category" binding:"required"`
		Description string  `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateExpense(expenseID, userID, input.Amount, input.Category, input.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
}

// DeleteExpenseHandler godoc
// @Summary      Delete expense
// @Description  Delete an expense by ID
// @Tags         expenses
// @Produce      json
// @Param        id   path      int  true  "Expense ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/expenses/{id} [delete]
func DeleteExpenseHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	idStr := c.Param("id")
	expenseID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	if err := service.DeleteExpense(expenseID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}

// GetFilteredExpensesHandler godoc
// @Summary      Get filtered expenses
// @Description  Get expenses by filter
// @Tags         expenses
// @Produce      json
// @Param        from     query   string  false  "From date"
// @Param        to       query   string  false  "To date"
// @Param        category query   string  false  "Category"
// @Param        min      query   number  false  "Min amount"
// @Param        max      query   number  false  "Max amount"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/expenses/categories [get]
func GetFilteredExpensesHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	from := c.Query("from")
	to := c.Query("to")
	category := c.Query("category")
	minStr := c.Query("min")
	maxStr := c.Query("max")

	var min, max float64
	var err error

	if minStr != "" {
		min, err = strconv.ParseFloat(minStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minimum amount"})
			return
		}
	}
	if maxStr != "" {
		max, err = strconv.ParseFloat(maxStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maximum amount"})
			return
		}
	}
	expenses, err := service.GetFilteredExpenses(userID, from, to, category, min, max)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve filtered expenses", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "message": "Filtered expenses retrieved successfully"})

}

// GetTotalFilteredExpensesHandler godoc
// @Summary      Get total filtered expenses
// @Description  Get total amount of filtered expenses
// @Tags         expenses
// @Produce      json
// @Param        from     query   string  false  "From date"
// @Param        to       query   string  false  "To date"
// @Param        category query   string  false  "Category"
// @Param        min      query   number  false  "Min amount"
// @Param        max      query   number  false  "Max amount"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/expenses/total [get]
func GetTotalFilteredExpensesHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	from := c.Query("from")
	to := c.Query("to")
	category := c.Query("category")
	minStr := c.Query("min")
	maxStr := c.Query("max")

	var min, max float64
	var err error

	if minStr != "" {
		min, err = strconv.ParseFloat(minStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minimum amount"})
			return
		}
	}

	if maxStr != "" {
		max, err = strconv.ParseFloat(maxStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maximum amount"})
			return
		}
	}
	total, err := service.GetTotalFilterExpenses(userID, from, to, category, min, max)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total filtered expenses", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "message": "Total filtered expenses calculated successfully"})

}
