package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudolfoborges/go-blog/config"
	"github.com/rudolfoborges/go-blog/ctx"
	"github.com/rudolfoborges/go-blog/models"
	"github.com/rudolfoborges/go-blog/usecases"
	"github.com/rudolfoborges/go-blog/utils"
)

func CreateUserHandler(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request body"})
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	err := usecases.CreateUserUseCase(&user)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user.Serialize())
}

func GetAllUsersHandler(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(400, gin.H{"message": "Failed to get users"})
		return
	}

	c.JSON(200, utils.ToSerializedArray(users))
}

func GetUserHandler(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"message": "Failed to get user"})
		return
	}

	c.JSON(200, user.Serialize())
}

func UpdatePasswordHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var body struct {
		CurrentPassword string `json:"currentPassword" binding:"required"`
		NewPassword     string `json:"newPassword" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request body"})
		return
	}

	ctx := ctx.UpdatePasswordContext{
		CurrentPassword: body.CurrentPassword,
		NewPassword:     body.NewPassword,
	}

	err := usecases.UpdatePasswordUseCase(id, ctx)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Password updated"})
}
