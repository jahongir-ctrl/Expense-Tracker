package repository

import (
	"ExpenceTracker/internal/db"
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/utils"
	"errors"
)

func CreateUser(fullName, username, password string) error {
	conn := db.GetDBConn()

	hashed := utils.HashPassword(password)
	query := `INSERT INTO users (full_name, username, password) VALUES ($1, $2, $3)`

	_, err := conn.Exec(query, fullName, username, hashed)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	conn := db.GetDBConn()
	var user models.User

	query := `SELECT * FROM users WHERE username = $1`

	err := conn.Get(&user, query, username)
	if err != nil {
		return user, errors.New("User not found")
	}
	return user, nil
}
