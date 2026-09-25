package domain

import "context"

type UserRepository interface {
	FindByEmail(
		ctx context.Context,
		email string,
	) (*User, error)

	FindByID(
		ctx context.Context,
		id uint64,
	) (*User, error)

	Create(
		ctx context.Context,
		user *User,
	) error
}