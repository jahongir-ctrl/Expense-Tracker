package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
)

func CreateGoal(goal models.Goal) error {
	query := `
INSERT INTO goals (user_id, title, target_amount, current_amount, description)
VALUES ($1, $2, $3, $4, $5)
`
	_, err := db.GetDBConn().Exec(query, goal.UserID, goal.Title, goal.TargetAmount, goal.CurrentAmount, goal.Description)
	return err
}

func GetGoalsByUser(userID int) ([]models.Goal, error) {
	var goals []models.Goal
	query := `SELECT * FROM goals WHERE user_id = $1 ORDER BY created_at DESC`
	err := db.GetDBConn().Select(&goals, query, userID)
	return goals, err
}

func DeleteGoal(id, userID int) error {
	query := `DELETE FROM goals WHERE id = $1 AND user_id = $2`
	_, err := db.GetDBConn().Exec(query, id, userID)
	return err
}

func UpdateGoalAmount(goalID, userID int, amount float64) error {
	query := `UPDATE goals SET current_amount = current_amount + $1 WHERE id = $2 AND user_id = $3`
	_, err := db.GetDBConn().Exec(query, amount, goalID, userID)
	return err
}
