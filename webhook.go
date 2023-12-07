package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func sendWebhookNotification(data []byte) error {
	webhookURLs := []string{
		"https://webhook.site/2898cf90-f0a5-4718-b5cd-9ab40fa2f498",
	}

	for _, url := range webhookURLs {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			return fmt.Errorf("failed to send webhook to %s: %w", url, err)
		}
		log.Printf("Event sent: %s\n", data)
		defer resp.Body.Close()
	}

	return nil
}
