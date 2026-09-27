package mongodb

import (
	"context"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDb(cfg *config.MongoConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(cfg.MongoUrl).
		SetAuth(options.Credential{
			Username: cfg.MongoUser,
			Password: cfg.MongoPwd,
			AuthSource: cfg.MongoName,
		})
	
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, err
	}

	return client, nil
}