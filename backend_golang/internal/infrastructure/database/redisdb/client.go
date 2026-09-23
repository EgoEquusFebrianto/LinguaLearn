package redisdb

import (
	"context"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisADDR,
		Password: cfg.RedisPwd,
		DB: cfg.RedisDb,
	})

	ctx := context.Background()

	if err :=client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, err
	}

	return client, nil
}