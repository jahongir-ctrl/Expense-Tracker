package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
)

func CreateGoal(goal models.Goal) error {
	return repository.CreateGoal(goal)
}

func GetGoalsByUser(userID int) ([]models.Goal, error) {
	return repository.GetGoalsByUser(userID)
}

func DeleteGoal(id, userID int) error {
	return repository.DeleteGoal(id, userID)
}

func UpdateGoalAmount(goalID, userID int, amount float64) error {
	return repository.UpdateGoalAmount(goalID, userID, amount)
}
