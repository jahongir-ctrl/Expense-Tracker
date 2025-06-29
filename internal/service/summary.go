package service

import "ExpenceTracker/internal/repository"

type BalanceSummary struct {
	Income   float64 `json:"total_income"`
	Expenses float64 `json:"total_expenses"`
	Balance  float64 `json:"balance"`
}

func GetBalanceSummary(userID int) (*BalanceSummary, error) {
	income, err := repository.GetTotalIncome(userID)
	if err != nil {
		return nil, err
	}

	expenses, err := repository.GetTotalExpense(userID)
	if err != nil {
		return nil, err
	}

	return &BalanceSummary{
		Income:   income,
		Expenses: expenses,
		Balance:  income - expenses,
	}, nil
}
