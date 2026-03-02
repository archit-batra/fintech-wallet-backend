package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/archit-batra/fintech-wallet-backend/internal/audit"
	"github.com/archit-batra/fintech-wallet-backend/internal/events"
	"github.com/archit-batra/fintech-wallet-backend/internal/infra"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Worker started...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=wallet sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	auditRepo := audit.NewRepository(db)
	redisClient := infra.NewRedisClient()
	defer redisClient.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker shutting down gracefully")
			return
		default:
			// Use timeout of 5 seconds instead of 0
			result, err := redisClient.BRPop(ctx, 5*time.Second, "transfer_queue").Result()

			if err != nil {
				// Timeout error is expected when queue is empty
				continue
			}

			var event events.TransferEvent
			if err := json.Unmarshal([]byte(result[1]), &event); err != nil {
				continue
			}

			auditRepo.InsertLog(
				event.EventType,
				event.FromUser,
				event.ToUser,
				event.Amount,
			)
		}
	}
}
