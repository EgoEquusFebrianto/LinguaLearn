package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/mongodb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/mysql"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/redis"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/route"
)

func Run() error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	db, err := mysql.NewMySQL(&cfg.MySQL)
	if err != nil {
		return err
	}
	defer db.Close()

	mongoClient, err := mongodb.NewMongoDb(&cfg.MongoDb)
	if err != nil {
		return err
	}
	defer func() {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			5 * time.Second,
		)
		defer cancel()

		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("failed to disconnect MongoDB: %v", err)
		}
	}()

	redisClient, err := redis.NewRedis(&cfg.Redis)
	if err != nil {
		return err
	}
	defer redisClient.Close()

	gormDB, err := mysql.NewGorm(db)
	if err != nil {
		return err
	}

	deps := &helper.Dependencies{
		MySQL: db,
		MongoDB: mongoClient,
		Redis: redisClient,
		GORM: gormDB,
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