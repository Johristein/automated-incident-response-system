package event

import (
	"fmt"
	"time"
)

// SecurityEvent represents a security-related incident
type SecurityEvent struct {
	ID              int
	Message         string
	IP              string
	OpenPorts       []int
	Software        string
	OperatingSystem string
	Details         string
}

func ProcessEvent(event SecurityEvent, result chan string) {
	fmt.Printf("Processing event ID %d: %s\n", event.ID, event.Message)

	// Simulate processing delay
	time.Sleep(2 * time.Second)

	result <- fmt.Sprintf("Event ID %d processed successfully!", event.ID)
}
