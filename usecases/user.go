package usecases

import (
	"errors"
	"log"

	"github.com/rudolfoborges/go-blog/config"
	"github.com/rudolfoborges/go-blog/ctx"
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

func UpdatePasswordUseCase(id int, ctx ctx.UpdatePasswordContext) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("User not found")
	}

	if err := user.ComparePassword(ctx.CurrentPassword); err != nil {
		return errors.New("Current password is incorrect")
	}

	user.UpdatePassword(ctx.NewPassword)
	config.DB.Save(&user)

	return nil
}
