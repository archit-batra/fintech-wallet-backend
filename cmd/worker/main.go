package main

import (
	"log"
	"time"

	"github.com/archit-batra/fintech-wallet-backend/internal/events"
)

func main() {

	log.Println("Worker started...")

	for {
		select {
		case event := <-events.EventQueue:
			log.Printf("Processing event: %s | %s", event.Type, event.Data)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
