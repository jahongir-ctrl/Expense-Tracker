package controller

import (
	"github.com/gin-gonic/gin"
	//"ExpenseTracker/internal/controller"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", RegisterHandler)
	r.POST("/login", LoginHandler)

	auth := r.Group("/api")
	auth.Use(AuthMiddleware())

	auth.GET("/profile", ProfileHandler)

	expenseGroup := auth.Group("/expenses")
	expenseGroup.POST("", CreateExpenseHandler)
	expenseGroup.GET("", GetExpencesHandler)
	expenseGroup.PUT("/:id", UpdateExpenseHandler)
	expenseGroup.DELETE("/:id", DeleteExpenseHandler)
	expenseGroup.GET("/categories", GetFilteredExpensesHandler)
	expenseGroup.GET("/total", GetTotalFilteredExpensesHandler)

	budgetGroup := auth.Group("/budgets")
	budgetGroup.POST("", SetBudgetHandler)
	budgetGroup.GET("", GetBudgetsHandler)
	budgetGroup.GET("/status", CheckBudgetHandler)

	incomeGroup := auth.Group("/incomes")
	incomeGroup.POST("", CreateIncomeHandler)
	incomeGroup.GET("", GetIncomesHandler)
	incomeGroup.DELETE("/:id", DeleteIncomeHandler)

	goalsGroup := auth.Group("/goals")
	goalsGroup.POST("", CreateGoalHandler)
	goalsGroup.GET("", GetGoalsHandler)
	goalsGroup.PUT("/:id", UpdateGoalHandler)
	goalsGroup.DELETE("/:id", DeleteGoalHandler)

	categoryGroup := auth.Group("/categories")
	categoryGroup.POST("", CreateCategoryHandler)
	categoryGroup.GET("", GetCategoriesHandler)
	categoryGroup.DELETE("/:id", DeleteCategoryHandler)

	auth.GET("/summary/balance", GetBalanceSummaryHandler)

	reportGroup := auth.Group("/reports")
	reportGroup.GET("/daily", GetDailyReportHandler)
	reportGroup.GET("/weekly", GetWeeklyReportHandler)
	reportGroup.GET("/monthly", GetMonthlyReportHandler)
}
