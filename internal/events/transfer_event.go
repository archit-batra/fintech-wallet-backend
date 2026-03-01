package events

import "time"

type TransferEvent struct {
	EventType string    `json:"event_type"`
	FromUser  int       `json:"from_user"`
	ToUser    int       `json:"to_user"`
	Amount    int64     `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
