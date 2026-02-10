package main

import (
	"github.com/gin-gonic/gin"

	"github.com/archit-batra/fintech-wallet-backend/internal/user"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	userService := user.NewService()
	userHandler := user.NewHandler(userService)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)

	router.Run(":8081")
}
