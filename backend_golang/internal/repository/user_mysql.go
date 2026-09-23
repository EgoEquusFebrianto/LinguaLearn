package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(
		ctx context.Context,
		email string,
	) (*domain.User, error)

	FindByID (
		ctx context.Context,
		id uint64,
	) (*domain.User, error)

	Create(
		ctx context.Context,
		user *domain.User,
	) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByEmail (
	ctx context.Context,
	email string,
) (*domain.User, error) {
	var user domain.User

	err := r.db.
		WithContext(ctx).
		Preload("Role").
		Where("email = ?", email).
		First(&user).
		Error 
		
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to find user by ID: %w", err)
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return &user, nil
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*domain.User, error) {
	var user domain.User

	err := r.db.
		WithContext(ctx).
		Preload("Role").
		First(&user, id).
		Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return &user, nil
}

func (r *userRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {
	return r.db.
		WithContext(ctx).
		Create(user).
		Error
}