package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create a test recorder and request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// 1. Check status code
	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, expectedStatus)
	}

	// 2. Check Content-Type header
	expectedContentType := "application/json"
	actualContentType := rr.Header().Get("Content-Type")
	if !strings.Contains(actualContentType, expectedContentType) {
		t.Errorf("handler returned wrong Content-Type header: got %v, want %s", actualContentType, expectedContentType)
	}

	// 3. Check body content
	var actualStatus map[string]string
	err = json.NewDecoder(rr.Body).Decode(&actualStatus)
	if err != nil {
		t.Fatalf("Could not decode response body JSON: %v", err)
	}

	expectedStatusMap := map[string]string{"status": "ok"}
	for key, expectedValue := range expectedStatusMap {
		if actualStatus[key] != expectedValue {
			t.Errorf("Status field mismatch for key %s: got %q, want %q", key, actualStatus[key], expectedValue)
		}
	}
}
