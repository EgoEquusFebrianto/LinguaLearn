package router

import (
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/handler"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/middleware"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils/deps"
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


	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", deps.AuthHandler.Register)
		r.Post("/login", deps.AuthHandler.Login)
		r.Post("/logout", deps.AuthHandler.Logout)
		r.Post("/refresh", deps.AuthHandler.Refresh)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(deps.JwtService))
		r.Get("/users/me", deps.AuthHandler.GetMe)

		r.Get("/dictionary", deps.DictionaryHandler.Search)
		r.Get("/dictionary/word", deps.DictionaryHandler.FindByWord)
	})

	return r
}