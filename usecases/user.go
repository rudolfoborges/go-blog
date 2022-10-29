package usecases

import (
	"errors"
	"log"

	"github.com/rudolfoborges/go-blog/config"
	"github.com/rudolfoborges/go-blog/models"
)

func CreateUserUseCase(user *models.User) error {
	if err := config.DB.Find(&user, "email = ?", user.Email); err == nil {
		return errors.New("User already exists")
	}

	if err := config.DB.Create(user); err != nil {
		log.Println("Failed to create user", err)
		return err.Error
	}

	return nil
}
