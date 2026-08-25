package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/health",
		nil,
	)

	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"Expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var response map[string]string
	
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["status"] != "ok" {
		t.Errorf(
			"Expected status 'ok', got %s",
			response["status"],
		)
	}
}