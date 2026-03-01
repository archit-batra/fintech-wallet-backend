package main

import (
	"log"

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

		log.Println("Processing job:", result[1])
	}
}
