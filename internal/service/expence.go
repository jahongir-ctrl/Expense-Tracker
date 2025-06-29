package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
	"fmt"
)

func CreateExpense(userID int, amount float64, description, category string) error {
	// Check if budget is set and would be exceeded
	_, total, limit, err := IsBudgetExceeded(userID, category)
	if err != nil {
		return err
	}
	if limit > 0 && (total+amount) > limit {
		return fmt.Errorf("budget limit exceeded for category '%s': limit=%.2f, current=%.2f, new=%.2f", category, limit, total, total+amount)
	}
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
