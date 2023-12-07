package main

import (
	"encoding/json"
	"log"
)

type Event struct {
	Event     string    `json:"event"`
	EventDate string    `json:"event_date"`
	Data      EventData `json:"data"`
}

type EventData struct {
	Type   string    `json:"type"`
	Covers CoverData `json:"covers"`
}

type CoverData struct {
	BikesCount string `json:"bikes_count"`
	BikeValue  string `json:"bike_value"`
}

func processEvent(event Event) error {
	log.Printf("Processing Event: %s\n", event.Event)

	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return sendWebhookNotification(jsonData)
}
