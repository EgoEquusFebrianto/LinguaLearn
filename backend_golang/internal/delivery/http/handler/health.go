package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string {
		"status": "ok",
		"messagee": "LinguaLearn API is running.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func DatabaseHealthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(r.Context()); err != nil {
			response := map[string]string {
				"status": "error",
				"message": "MySQL database is unavailable.",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)

			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]string {
			"status": "ok",
			"message": "MySQL database is healthy",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(response)
	}
}