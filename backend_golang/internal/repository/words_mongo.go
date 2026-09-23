package repository

import (
	"context"
	"errors"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ErrBankWordNotFound = errors.New("Bank word not found.")

type BankWordRepository interface {
	FindByWord(
		ctx context.Context,
		word string,
	) (*domain.BankWord, error)

	FindByUUID(
		ctx context.Context,
		wordUUID string,
	) (*domain.BankWord, error)

	Search(
		ctx context.Context,
		query string,
		limit int,
		offset int,
	) ([]domain.BankWord, error)

	Count(
		ctx context.Context,
		query string,
	) (int64, error)
}

type bankWordRepository struct {
	collection *mongo.Collection
}

func NewBankWordRepository(
	client *mongo.Client,
	databaseName string,
) BankWordRepository {
	return &bankWordRepository{
		collection: client.
			Database(databaseName).
			Collection("bank_words"),
	}
}

func (r *bankWordRepository) FindByWord(
	ctx context.Context,
	word string,
) (*domain.BankWord, error) {
	var result domain.BankWord

	err := r.collection.
		FindOne(
			ctx,
			bson.M{
				"word": word,
			},
		).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrBankWordNotFound
		}

		return  nil, err
	}

	return &result, nil
}

func (r *bankWordRepository) FindByUUID(
	ctx context.Context,
	wordUUID string,
) (*domain.BankWord, error) {
	var result domain.BankWord

	err := r.collection.
		FindOne(
			ctx,
			bson.M{
				"word_uuid": wordUUID,
			},
		).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrBankWordNotFound
		}

		return  nil, err
	}

	return &result, nil
}

func (r *bankWordRepository) Search(
	ctx context.Context,
	query string,
	limit int,
	offset int,
) ([]domain.BankWord, error) {
	filter := bson.M{}

	if query != "" {
		filter["word"] = bson.M{
			"$regex": query,
			"$options": "i",
		}
	}

	findOpts  := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.D{
			{Key: "word", Value: 1},
		})

	cursor, err := r.collection.Find(
		ctx,
		filter,
		findOpts,
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	result := make([]domain.BankWord, 0)
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *bankWordRepository) Count(
	ctx context.Context,
	query string,
) (int64, error) {
	filter := bson.M{}

	if query != "" {
		filter["word"] = bson.M{
			"$regex": query,
			"$options": "i",
		}
	}

	return r.collection.CountDocuments(
		ctx,
		filter,
	)
}