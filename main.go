package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	go mockEventGenerator()
	http.HandleFunc("/receive-event", receiveEvent)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {

	var mockEvent Event
	if err := json.NewDecoder(r.Body).Decode(&mockEvent); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing event: %s", err), http.StatusBadRequest)
		return
	}

	if err := processEvent(mockEvent); err != nil {
		http.Error(w, fmt.Sprintf("Error processing event: %s", err), http.StatusInternalServerError)
		return
	} else {
		log.Println("Event received and processed")
	}

}
