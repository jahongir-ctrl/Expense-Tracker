package repository

import "ExpenceTracker/internal/db"

func GetTotalIncome(userID int) (float64, error) {
	var total float64
	query := `SELECT COALESCE(SUM(amount), 0) FROM incomes WHERE user_id = $1`
	err := db.GetDBConn().Get(&total, query, userID)
	return total, err
}

func GetTotalExpense(userID int) (float64, error) {
	var total float64
	query := `SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id = $1`
	err := db.GetDBConn().Get(&total, query, userID)
	return total, err
}
