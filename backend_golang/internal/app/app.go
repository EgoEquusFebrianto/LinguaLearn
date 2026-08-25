package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/route"
)

func Run() error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	db, err := database.NewMySQL(&cfg.MySQL)
	if err != nil {
		return err
	}
	defer db.Close()

	mongoClient, err := database.NewMongoDb(&cfg.MongoDb)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5 * time.Second,
	) 
	defer cancel()
	mongoClient.Disconnect(ctx)

	redisClient, err := database.NewRedis(&cfg.Redis)
	if err != nil {
		return err
	}
	defer redisClient.Close()

	deps := &helper.Dependencies{
		MySQL: db,
		MongoDB: mongoClient,
		Redis: redisClient,
	}

	router := route.NewRouter(deps)

	server := &http.Server{
		Addr: cfg.Server.Address,
		Handler: router,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Printf(
		"LinguaLearn API running on %s",
		cfg.Server.Address,
	)

	return server.ListenAndServe()
}