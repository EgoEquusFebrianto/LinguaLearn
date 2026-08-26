package main

import (
	"context"
	"log"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/database/mongodb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/migration/mongodb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/seed"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// MongoDB
	client, err := mongodb.NewMongoDb(&cfg.MongoDb)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB ", err)
	}

	// MongoDB Migration
	if err := migration.MigrateMongoDB(client, cfg.MongoDb.MongoDb); err != nil {
		log.Fatal("MongoDB migration failed: ", err)
	}

	log.Printf("MongoDB Migration Successfully.")

	// Dictionary Seed
	if err := seed.RunDictionaryseed(
		client,
		cfg.MongoDb.MongoDb,
		"seeds/mongodb/dictionary.json",
	); err != nil {
		log.Fatal("Dictionary seed failed: ", err)
	}

	log.Println("bank_words seed successfully.")

	ctx, cancel := context.WithTimeout(
		context.Background(), 
		30 * time.Second,
	)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
	}
}