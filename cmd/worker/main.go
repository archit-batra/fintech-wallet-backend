package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/archit-batra/fintech-wallet-backend/internal/audit"
	"github.com/archit-batra/fintech-wallet-backend/internal/events"
	"github.com/archit-batra/fintech-wallet-backend/internal/infra"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Worker started...")

	// DB connection
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=wallet sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	auditRepo := audit.NewRepository(db)
	redisClient := infra.NewRedisClient()
	ctx := context.Background()

	for {

		result, err := redisClient.BRPop(ctx, 0, "transfer_queue").Result()
		if err != nil {
			log.Println("Redis error:", err)
			continue
		}

		var event events.TransferEvent

		err = json.Unmarshal([]byte(result[1]), &event)
		if err != nil {
			log.Println("Invalid event payload:", err)
			continue
		}

		err = auditRepo.InsertLog(
			event.EventType,
			event.FromUser,
			event.ToUser,
			event.Amount,
		)

		if err != nil {
			log.Println("Failed to write audit log:", err)
			continue
		}

		log.Printf(
			"Audit log inserted | From: %d | To: %d | Amount: %d",
			event.FromUser,
			event.ToUser,
			event.Amount,
		)
	}
}
