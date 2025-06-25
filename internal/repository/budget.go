package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
)

func SetBudget(userID int, category string, limit float64) error {

	query := `INSERT INTO budgets (user_id, category, limit_amount) VALUES ($1, $2, $3)`
	_, err := db.GetDBConn().Exec(query, userID, category, limit)
	return err
}
func GetBudgets(userID int) ([]models.Budget, error) {
	var budgets []models.Budget
	query := `SELECT * FROM budgets WHERE user_id = $1`
	err := db.GetDBConn().Select(&budgets, query, userID)
	return budgets, err
}

func GetCategoryTotal(userID int, category string) (float64, error) {
	var total float64
	query := `SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id = $1 AND category = $2`
	err := db.GetDBConn().Get(&total, query, userID, category)
	return total, err
}

func GetBudgetLimit(userID int, category string) (float64, error) {
	query := `SELECT limit_amount FROM budgets WHERE user_id = $1 AND category = $2`
	var limit float64
	err := db.GetDBConn().Get(&limit, query, userID, category)
	return limit, err
}
