package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
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

// validateEvent checks if the event data is valid
func validateEvent(event Event) error {

	if strings.TrimSpace(event.Event) == "" {
		return errors.New("event type is missing or empty")
	}

	if err := validateEventDate(event.EventDate); err != nil {
		return err
	}

	return validateEventData(event.Data)
}

func validateEventDate(date string) error {
	if _, err := time.Parse(time.RFC3339, date); err != nil {
		return errors.New("invalid event date format")
	}
	return nil
}

func validateEventData(data EventData) error {

	if strings.TrimSpace(data.Type) == "" {
		return errors.New("event data type is missing or empty")
	}

	return validateCoverData(data.Covers)
}

func validateCoverData(cover CoverData) error {
	if count, err := strconv.Atoi(cover.BikesCount); err != nil || count <= 0 {
		return errors.New("invalid bikes count")
	}

	if value, err := strconv.Atoi(cover.BikeValue); err != nil || value < 0 {
		return errors.New("invalid bike value")
	}

	return nil
}

func processEvent(event Event) error {
	log.Printf("Processing Event: %s\n", event.Event)

	if err := validateEvent(event); err != nil {
		log.Printf("Validation error: %s\n", err)
		return err
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return sendWebhookNotification(jsonData)
}
