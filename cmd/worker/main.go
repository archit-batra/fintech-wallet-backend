package main

import (
	"encoding/json"
	"log"

	"github.com/archit-batra/fintech-wallet-backend/internal/events"
	"github.com/archit-batra/fintech-wallet-backend/internal/infra"
)

func main() {

	log.Println("Worker started...")

	redisClient := infra.NewRedisClient()

	for {
		result, err := redisClient.BRPop(
			infra.Ctx,
			0,
			"transfer_queue",
		).Result()

		if err != nil {
			log.Println("Error receiving job:", err)
			continue
		}

		var event events.TransferEvent

		err = json.Unmarshal([]byte(result[1]), &event)
		if err != nil {
			log.Println("Invalid event payload:", err)
			continue
		}

		log.Printf(
			"Transfer processed | From: %d | To: %d | Amount: %d | Time: %s",
			event.FromUser,
			event.ToUser,
			event.Amount,
			event.Timestamp,
		)
	}
}
