package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
)

func CreateRecurringExpense(exp models.Recurringexpense) error {
	return repository.CreateRecurringExpense(exp)
}

func DeleteRecurringExpense(id, userID int) error {
	return repository.DeleteRecurringExpense(id, userID)
}

func StopRecurringExpense(id, userID int) error {
	return repository.StopRecurringExpense(id, userID)
}

func GetRecurringExpenses(userID int) ([]models.Recurringexpense, error) {
	return repository.GetRecurringExpenses(userID)
}
