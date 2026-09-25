package domain

import "context"

type BankWordRepository interface {
	FindByWord(
		ctx context.Context,
		word string,
	) (*BankWord, error)

	FindByUUID(
		ctx context.Context,
		wordUUID string,
	) (*BankWord, error)

	Search(
		ctx context.Context,
		query string,
		limit int,
		offset int,
	) ([]BankWord, error)

	Count(
		ctx context.Context,
		query string,
	) (int64, error)
}