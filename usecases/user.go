package usecases

import (
	"errors"
	"log"

	"github.com/rudolfoborges/go-blog/config"
	"github.com/rudolfoborges/go-blog/models"
)

func CreateUserUseCase(user *models.User) error {
	var alreadyExists models.User
	if result := config.DB.Find(&alreadyExists, "email = ?", user.Email); result.RowsAffected > 0 {
		return errors.New("User already exists")
	}

	user.Role = "user"

	if err := config.DB.Create(user); err != nil {
		log.Println("Failed to create user", err)
		return err.Error
	}

	return nil
}
