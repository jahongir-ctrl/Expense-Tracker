package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
)

func CreateCategory(category models.Category) error {
	query := `INSERT INTO categories (user_id, name, type) VALUES ($1, $2, $3)`
	_, err := db.GetDBConn().Exec(query, category.UserID, category.Name, category.Type)
	return err
}

func GetCategoriesByUser(userID int) ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM categories WHERE user_id = $1 ORDER BY id DESC`
	err := db.GetDBConn().Select(&categories, query, userID)
	return categories, err
}

func DeleteCategory(categoryID, userID int) error {
	query := `DELETE FROM categories WHERE id = $1 AND user_id = $2`
	_, err := db.GetDBConn().Exec(query, categoryID, userID)
	return err
}
