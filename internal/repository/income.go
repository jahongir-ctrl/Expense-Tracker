package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
)

func CreateIncome(income models.Income) error {
	query := `
INSERT INTO incomes (user_id, amount, source, description)
VALUES ($1, $2, $3, $4)
`
	_, err := db.GetDBConn().Exec(query, income.UserID, income.Amount, income.Source, income.Description)
	return err
}

func GetIncomesByUser(userID int) ([]models.Income, error) {
	var incomes []models.Income
	query := `SELECT * FROM incomes WHERE user_id = $1 ORDER BY date DESC`
	err := db.GetDBConn().Select(&incomes, query, userID)
	return incomes, err
}

func UpdateIncome(income models.Income) error {
	query := `
UPDATE incomes SET amount = $1, source = $2, description = $3, date = $4 WHERE id = $5 AND user_id = $6
`
	_, err := db.GetDBConn().Exec(query, income.Amount, income.Source, income.Description, income.Date, income.ID, income.UserID)
	return err
}

func DeleteIncome(id, userID int) error {
	query := `
DELETE FROM incomes WHERE id = $1 AND user_id = $2
`
	_, err := db.GetDBConn().Exec(query, id, userID)
	return err
}
