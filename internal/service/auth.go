package service

import (
	"ExpenceTracker/internal/repository"
	"ExpenceTracker/internal/utils"
)

func RegisterUser(fullName, username, password string) error {
	return repository.CreateUser(fullName, username, password)
}
func LoginUser(userName, password string) (string, error) {
	user, err := repository.GetUserByUsername(userName)
	if err != nil {
		return "", err
	}

	hashed := utils.GenerateHash(password)
	if hashed != user.Password {
		return "", err
	}

	return utils.GenerateToken(user.ID, user.Username)
}
