package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

// HealthStatus represents the structure for the health check response.
type HealthStatus struct {
	Status string `json:"status"`
}

// healthHandler writes the structured JSON response for the /health endpoint.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set the required Content-Type header
	w.Header().Set("Content-Type", "application/json")
	
	// Create the response body
	status := HealthStatus{
		Status: "ok",
	}

	// Encode the structure to JSON and write to the response writer
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Printf("Error encoding status JSON: %v", err)
		// Fallback in case of encoding error
		http.Error(w, `{"status":"error", "message":"internal encoding error"}`, http.StatusInternalServerError)
	}
}

func main() {
	// Register the handler for the /health endpoint
	http.HandleFunc("/health", healthHandler)

	// Initialize and start the HTTP server
	addr := fmt.Sprintf(":%s", PORT)
	log.Printf("Starting HTTP server on http://localhost%s", addr)
	
	// Start listening on the port
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}