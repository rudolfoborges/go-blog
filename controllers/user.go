package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rudolfoborges/go-blog/models"
	"github.com/rudolfoborges/go-blog/usecases"
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

	c.JSON(200, gin.H{
		"message": "User created",
		"user_id": user.ID,
	})
}
