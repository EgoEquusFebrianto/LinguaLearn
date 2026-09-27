package main

import (
	"context"
	"log"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/database/scripts/migration"
	"github.com/EgoEquusFebrianto/LinguaLearn/database/scripts/seed"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/infrastructure/database/mongodb"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/infrastructure/database/mysql"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//MySQL
	db, err := mysql.NewMySQL(&cfg.MySQL)
	if err != nil {
		log.Fatal("Failed to connect to MySQL: ", err)
	}
	defer db.Close()

	if err := seed.SeedMySQL(db); err != nil {
		log.Fatal("Failed to seed MySQL: ", err)
	}
	log.Println("MySQL seed completed successfully.")

	// MongoDB
	client, err := mongodb.NewMongoDb(&cfg.MongoDb)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB ", err)
	}

	// MongoDB Migration
	if err := migration.MigrateMongoDB(client, cfg.MongoDb.MongoName); err != nil {
		log.Fatal("MongoDB migration failed: ", err)
	}

	log.Printf("MongoDB Migration Successfully.")

	// Dictionary Seed
	if err := seed.RunDictionaryseed(
		client,
		cfg.MongoDb.MongoName,
		"database/seeds/mongodb/dictionary.json",
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