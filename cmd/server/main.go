package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/archit-batra/fintech-wallet-backend/internal/user"
	"github.com/archit-batra/fintech-wallet-backend/internal/wallet"
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

	walletRepo := wallet.NewRepository(db)
	walletService := wallet.NewService(walletRepo)
	walletHandler := wallet.NewHandler(walletService)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)

	router.POST("/wallets/:userId", walletHandler.CreateWallet)
	router.POST("/wallets/:userId/add", walletHandler.AddMoney)
	router.GET("/wallets/:userId", walletHandler.GetWallet)

	router.Run(":8081")
}
