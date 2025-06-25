package controller

import (
	"github.com/gin-gonic/gin"
	//"ExpenceTracker/internal/controller"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", RegisterHandler)
	r.POST("/login", LoginHandler)

	auth := r.Group("/api")
	auth.Use(AuthMiddleware())

	auth.GET("/profile", ProfileHandler)
	auth.POST("/expenses", CreateExpenseHandler)
	auth.GET("/expenses", GetExpencesHandler)
	auth.PUT("/expenses/:id", UpdateExpenseHandler)
	auth.DELETE("/expenses/:id", DeleteExpenseHandler)
	auth.GET("/expenses/categories", GetFilteredExpensesHandler)
	auth.GET("/expenses/total", GetTotalFilteredExpensesHandler)

	auth.POST("/budgets", SetBudgetHandler)
	auth.GET("/budgets", GetBudgetsHandler)
	auth.GET("/budgets/status", CheckBudgetHandler)
}
