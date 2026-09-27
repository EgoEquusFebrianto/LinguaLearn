package mongodb

import (
	"context"
	"os"
	"testing"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils"
)

func TestNewMongoDB(t *testing.T) {
	utils.LoadEnv(t)

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