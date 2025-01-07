package main

import (
	"log"
	"net/http"
	"os"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a response recorder to capture the response
		recorder := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)
		// Log the response content
		log.Printf("Response: %s", recorder.body)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body = append(rec.body, b...)
	return rec.ResponseWriter.Write(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("./index.html")
	if err != nil {
		log.Printf("Error reading HTML file: %v", err)
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

func main() {
	// Open a file for logging
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set the output of the log package to the file
	log.SetOutput(logFile)

	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
