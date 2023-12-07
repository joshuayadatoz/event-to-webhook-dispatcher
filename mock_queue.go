package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func mockEventGenerator() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		<-ticker.C
		sendMockEvent()
	}

}

func sendMockEvent() {

	bikeCount := rand.Intn(9) + 1
	bikeValue := bikeCount * 140000
	mockEvent := Event{
		Event:     "policy.updated",
		EventDate: time.Now().Format(time.RFC3339),
		Data: EventData{
			Type: "adjustment",
			Covers: CoverData{
				BikesCount: fmt.Sprintf("%d", bikeCount),
				BikeValue:  fmt.Sprintf("%d", bikeValue),
			},
		},
	}

	eventData, err := json.Marshal(mockEvent)
	if err != nil {
		log.Println("Error marshalling mock event:", err)
		return
	}

	_, err = http.Post("http://localhost:8080/receive-event", "application/json", bytes.NewBuffer(eventData))
	if err != nil {
		log.Println("Error sending mock event:", err)
	}
}
