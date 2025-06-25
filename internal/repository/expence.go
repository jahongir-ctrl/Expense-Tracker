package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
	"fmt"
)

func CreateExpense(userID int, amount float64, category, description string) error {
	query := `INSERT INTO expenses (user_id, amount, category, description) VALUES ($1, $2, $3, $4)`
	_, err := db.GetDBConn().Exec(query, userID, amount, category, description)
	return err
}

func GetExpensesByUser(userID int) ([]models.Expense, error) {
	var expenses []models.Expense
	query := `SELECT * FROM expenses WHERE user_id = $1 ORDER BY created_at DESC`
	err := db.GetDBConn().Select(&expenses, query, userID)
	return expenses, err
}

func UpdateExpense(expenseID int, userID int, amount float64, category, description string) error {
	query := `UPDATE expenses SET amount = $1, category = $2, description = $3 WHERE id = $4 AND user_id = $5`
	_, err := db.GetDBConn().Exec(query, amount, category, description, expenseID, userID)
	return err
}

func DeleteExpense(expenseID, userID int) error {
	query := `DELETE FROM expenses WHERE id = $1 AND user_id = $2`
	_, err := db.GetDBConn().Exec(query, expenseID, userID)
	return err
}

func GetFilteredExpenses(userID int, from, to, category string, min, max float64) ([]models.Expense, error) {
	query := `SELECT *FROM expenses WHERE user_id = $1`
	args := []interface{}{userID}
	i := 2

	if from != "" {
		query += fmt.Sprintf(" AND created_at >= $%d", i)
		args = append(args, from)
		i++
	}

	if to != "" {
		query += fmt.Sprintf(" AND created_at <= $%d", i)
		args = append(args, to)
		i++
	}

	if category != "" {
		query += fmt.Sprintf(" AND category = $%d", i)
		args = append(args, category)
		i++
	}
	if min > 0 {
		query += fmt.Sprintf(" AND amount >= $%d", i)
		args = append(args, min)
		i++
	}

	if max > 0 {
		query += fmt.Sprintf(" AND amount <= $%d", i)
		args = append(args, max)
		i++
	}

	var expenses []models.Expense
	err := db.GetDBConn().Select(&expenses, query, args...)
	return expenses, err

}

func GetTotalFilterExpenses(userID int, from, to, category string, min, max float64) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id = $1`
	args := []interface{}{userID}
	i := 2

	if from != "" {
		query += fmt.Sprintf("%s AND created_at >= $%d", i)
		args = append(args, from)
		i++
	}

	if to != "" {
		query += fmt.Sprintf("%s AND created_at <= $%d", i)
		args = append(args, to)
		i++
	}

	if category != "" {
		query += fmt.Sprintf("%s AND category = $%d", i)
		args = append(args, category)
		i++
	}

	if min > 0 {
		query += fmt.Sprintf("%s AND amount >= $%d", i)
		args = append(args, min)
		i++
	}

	if max > 0 {
		query += fmt.Sprintf("%s AND amount <= $%d", i)
		args = append(args, max)
		i++
	}
	var total float64
	err := db.GetDBConn().Get(&total, query, args...)
	return total, err

}
