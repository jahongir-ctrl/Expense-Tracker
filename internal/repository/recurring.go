package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
)

func CreateRecurringExpense(exp models.Recurringexpense) error {
	query := `INSERT INTO recurring_expenses (user_id, amount, category, description, frequency, next_date, active)
VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.GetDBConn().Exec(query, exp.UserID, exp.Amount, exp.Category, exp.Description, exp.Frequency, exp.NextDate, exp.Active)
	return err
}

func GetRecurringExpenses(userID int) ([]models.Recurringexpense, error) {
	var result []models.Recurringexpense
	query := `SELECT * FROM recurring_expenses WHERE user_id = $1 ORDER BY created_at DESC`
	err := db.GetDBConn().Select(&result, query, userID)
	return result, err
}

func DeleteRecurringExpense(id, userID int) error {
	query := `DELETE FROM recurring_expenses WHERE id = $1 AND user_id = $2`
	_, err := db.GetDBConn().Exec(query, id, userID)
	return err
}

func StopRecurringExpense(id, userID int) error {
	query := `UPDATE recurring_expenses SET active = false WHERE id = $1 AND user_id = $2`
	_, err := db.GetDBConn().Exec(query, id, userID)
	return err
}
