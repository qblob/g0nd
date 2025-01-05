package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var eventChannel = make(chan string)

func handler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("./index.html")
	if err != nil {
		log.Printf("Error reading HTML file: %v", err)
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)

	// Send an event to the channel
	eventChannel <- "Page accessed"
}

func eventListener() {
	for {
		select {
		case event := <-eventChannel:
			log.Printf("Event received: %s", event)
		case <-time.After(10 * time.Second):
			log.Println("No events received in the last 10 seconds")
		}
	}
}

func main() {
	go eventListener()
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
