package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/deps"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/route"

)

type App struct {
	Config 	*config.Config
	Deps 	*deps.Dependencies
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	dependencies, err := SetupDependencies(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Deps: dependencies,
	}, nil
}

func (a *App) Run() error {
	router := route.NewRouter(a.Deps)

	server := &http.Server{
		Addr: a.Config.Server.Address,
		Handler: router,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Printf(
		"LinguaLearn API Running on %s",
		a.Config.Server.Address, 
	)

	return server.ListenAndServe()
}

func (a *App) Close() {
	if a.Deps.MySQL != nil {
		a.Deps.MySQL.Close()
	}

	if a.Deps.Redis != nil {
		a.Deps.Redis.Close()
	}

	if a.Deps.MongoDB != nil {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			5 * time.Second,
		)
		defer cancel()

		if err := a.Deps.MongoDB.Disconnect(ctx); err != nil {
			log.Printf(
				"Failed to disconnect MongoDB: %v",
				err,
			)
		}
	}
}

func Run() error {
	app, err := New()
	if err != nil {
		return err
	}

	defer app.Close()

	return app.Run()
}