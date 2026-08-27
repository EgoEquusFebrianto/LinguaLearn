package handler

import (
	"net/http"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/middleware"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		helper.Error(
			w,
			http.StatusUnauthorized,
			"User not found in context",
		)
		return
	}

	role, _ := middleware.GetRole(r.Context())

	response := map[string]interface{} {
		"user_id": userID,
		"role": role,
	}

	helper.JSON(w, http.StatusOK, response)
}