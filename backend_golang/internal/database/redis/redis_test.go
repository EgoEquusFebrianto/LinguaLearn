package redis

import (
	"os"
	"strconv"
	"testing"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/joho/godotenv"
)

func TestNewRedis(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Failed to load .env: %v", err)
	}

	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	cfg := &config.RedisConfig{
		RedisADDR: os.Getenv("REDIS_ADDR"),
		RedisPwd: os.Getenv("REDIS_PASSWORD"),
		RedisDb: redisDb,
	}

	client, err := NewRedis(cfg)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatal("Expected client, Not Nil")
	}

	defer client.Close()
}