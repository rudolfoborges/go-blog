package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rudolfoborges/go-blog/config"
	"github.com/rudolfoborges/go-blog/controllers"
)

func init() {
	godotenv.Load()
	config.InitDatabase()
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/users", controllers.CreateUserHandler)
	}

	r.Run(":3333")
}
