package events

// Event represents a background job
type Event struct {
	Type string
	Data string
}

// EventQueue is a simple channel-based queue
var EventQueue = make(chan Event, 100)
