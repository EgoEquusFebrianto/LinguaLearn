package route

import (
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/deps"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(deps *deps.Dependencies) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/v1/health", handler.HealthHandler)
	r.Get("/api/v1/health/database", handler.DatabaseHealthHandler(deps.MySQL))

	authHandler := handler.NewAuthHandler(deps.AuthService)

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(deps.JwtService))
		r.Get("/users/me", handler.GetMe)
	})

	return r
}