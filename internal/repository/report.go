package repository

import (
	"ExpenceTracker/internal/db"
	"time"
)

func GetExpenseSumDate(userID int, from, to time.Time) (float64, error) {
	var total float64
	query := `
		SELECT coalesce(SUM(amount), 0) FROM expenses WHERE user_id = $1 AND date BETWEEN $2 AND $3`
	err := db.GetDBConn().Get(&total, query, userID, from, to)
	return total, err
}
