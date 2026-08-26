package migration

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const bankWordsCollection = "bank_words"

func MigrateMongoDB(client *mongo.Client, databaseName string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30 * time.Second,
	)
	defer cancel()

	db := client.Database(databaseName)

	// Guarantee bank_words Collections available,
	collections, err := db.ListCollectionNames(ctx, bson.D{
		{Key: "name", Value: bankWordsCollection},
	})
	if err != nil {
		return fmt.Errorf("Failed to check bank_words collection: %w", err)
	}

	if len(collections) == 0 {
		if err := db.CreateCollection(ctx, bankWordsCollection); err != nil {
			return fmt.Errorf("Failed to create bank_words collection: %w", err)
		}
	}

	collection := db.Collection(bankWordsCollection)

	// unique index for word_uuid
	_, err = collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "word_uuid", Value: 1},
			},
			Options: options.Index().
				SetUnique(true).
				SetName("uq_bank_words_word_uuid"),
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to create word_uuid index: %w", err)
	}

	// Index for search word
	_, err = collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "word", Value: 1},		
			},
			Options: options.Index().
				SetName("idx_bank_words_word"),
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to create word Index: %w", err)
	}

	// Index for translations
	_, err = collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "sense.translations", Value: 1},
			},
			Options: options.Index().
				SetName("idx_bank_words_translations"),
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to create translations index: %w", err)
	}

	return nil
}