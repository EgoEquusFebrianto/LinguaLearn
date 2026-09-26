package redisdb

import (
	"os"
	"strconv"
	"testing"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils"
)

func TestNewRedis(t *testing.T) {
	utils.LoadEnv(t)

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