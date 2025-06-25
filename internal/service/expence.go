package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
)

func CreateExpense(userID int, amount float64, description, category string) error {
	return repository.CreateExpense(userID, amount, category, description)
}

func GetUserExpenses(userID int) ([]models.Expense, error) {
	return repository.GetExpensesByUser(userID)
}

func UpdateExpense(expenseID int, userID int, amount float64, category, description string) error {
	return repository.UpdateExpense(expenseID, userID, amount, category, description)
}

func DeleteExpense(expenseID, userID int) error {
	return repository.DeleteExpense(expenseID, userID)
}

func GetFilteredExpenses(userID int, from, to, category string, min, max float64) ([]models.Expense, error) {
	return repository.GetFilteredExpenses(userID, from, to, category, min, max)
}

func GetTotalFilterExpenses(userID int, from, to, category string, min, max float64) (float64, error) {
	return repository.GetTotalFilterExpenses(userID, from, to, category, min, max)
}
