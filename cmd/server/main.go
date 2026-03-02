package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/archit-batra/fintech-wallet-backend/internal/user"
	"github.com/archit-batra/fintech-wallet-backend/internal/wallet"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// DB connection
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
	router.POST("/wallets/transfer", walletHandler.Transfer)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		log.Println("Server running on :8081")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed:", err)
	}

	db.Close()

	log.Println("Server exited properly")
}
