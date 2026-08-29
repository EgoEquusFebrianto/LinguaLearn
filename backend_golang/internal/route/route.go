package route

import (
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/deps"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(deps *deps.Dependencies) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	r.Get("/api/v1/health", handler.HealthHandler)
	r.Get("/api/v1/health/database", handler.DatabaseHealthHandler(deps.MySQL))

	authHandler := handler.NewAuthHandler(deps.AuthService)

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/logout", authHandler.Logout)
		r.Post("/refresh", authHandler.Refresh)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(deps.JwtService))
		r.Get("/users/me", handler.GetMe)
	})

	return r
}