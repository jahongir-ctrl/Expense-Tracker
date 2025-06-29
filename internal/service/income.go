package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
	//"ExpenceTracker/internal/repository"
)

func CreateIncome(userID int, amount float64, description, source string, date string) error {
	income := models.Income{
		UserID:      userID,
		Amount:      amount,
		Description: description,
		Source:      source,
	}

	return repository.CreateIncome(income)
}

func GetIncomes(userID int) ([]models.Income, error) {
	return repository.GetIncomesByUser(userID)
}

func UpdateIncome(income models.Income) error {
	return repository.UpdateIncome(income)
}

func DeleteIncome(incomeID, userID int) error {
	return repository.DeleteIncome(incomeID, userID)
}
