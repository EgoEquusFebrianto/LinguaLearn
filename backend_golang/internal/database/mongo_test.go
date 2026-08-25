package database

import (
	"context"
	"os"
	"testing"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/joho/godotenv"
)

func TestNewMongoDB(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Failed to load .env: %v", err)
	}

	cfg := &config.MongoConfig{
		MongoUrl: os.Getenv("MONGO_URL"),
		MongoUser: os.Getenv("MONGO_USER"),
		MongoPwd: os.Getenv("MONGO_PASSWORD"),
		MongoDb: os.Getenv("MONGO_NAME"),
	}

	client, err := NewMongoDb(cfg)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatal("Expected Client not nil")
	}

	defer client.Disconnect(context.Background())
}