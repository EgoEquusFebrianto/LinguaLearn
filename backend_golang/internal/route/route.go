package route

import (
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/go-chi/chi/v5"
)

func NewRouter(deps *helper.Dependencies) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/v1/health", handler.HealthHandler)
	r.Get("/api/v1/health/database", handler.DatabaseHealthHandler(deps.MySQL))

	return r
}