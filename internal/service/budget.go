package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
	//"ExpenceTracker/internal/models"
)

func SetBudget(userID int, category string, limit float64) error {
	return repository.SetBudget(userID, category, limit)
}

func GetBudgets(userID int) ([]models.Budget, error) {
	return repository.GetBudgets(userID)
}

func IsBudgetExceeded(userID int, category string) (bool, float64, float64, error) {
	total, err := repository.GetCategoryTotal(userID, category)
	if err != nil {
		return false, 0, 0, err
	}
	limit, err := repository.GetBudgetLimit(userID, category)
	if err != nil {
		return false, 0, 0, err
	}
	return total > limit, total, limit, nil
}
