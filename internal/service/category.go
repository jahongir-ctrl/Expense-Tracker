package service

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/repository"
)

func CreateCategory(category models.Category) error {
	return repository.CreateCategory(category)
}

func GetCategoriesByUser(userID int) ([]models.Category, error) {
	return repository.GetCategoriesByUser(userID)
}

func DeleteCategory(id, userID int) error {
	return repository.DeleteCategory(id, userID)
}
