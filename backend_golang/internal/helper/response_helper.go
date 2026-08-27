package helper

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(
	w http.ResponseWriter,
	status int,
	data any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func Error(
	w http.ResponseWriter,
	status int,
	message string,
) {
	JSON(
		w,
		status,
		ErrorResponse{
			Error: message,
		},
	)
}