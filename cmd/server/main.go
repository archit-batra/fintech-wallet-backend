package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/archit-batra/fintech-wallet-backend/internal/user"
)

func main() {

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=wallet sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)

	router.Run(":8081")
}
