package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"github.com/redis/go-redis/v9"
)

const refreshTokenTTL = 7 * 24 * time.Hour

type RefreshTokenService struct {
	redis *redis.Client
}

func NewRefreshTokenService(
	redis *redis.Client,
) *RefreshTokenService {
	return &RefreshTokenService{
		redis: redis,
	}
}

func (s *RefreshTokenService) Create(
	ctx context.Context,
	UserID uint64,
) (string, error) {
	token, tokenHash, err := security.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	key := "refresh:" + tokenHash
	
	err = s.redis.Set(
		ctx,
		key,
		strconv.FormatUint(UserID, 10),
		refreshTokenTTL,
	).Err()

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *RefreshTokenService) GetUserID(
	ctx context.Context,
	token string,
) (uint64, error) {
	if token == "" {
		return 0, errors.New("Refresh token is required.")
	}

	tokenHash := security.HashRefreshToken(token)
	key := "refresh:" + tokenHash

	value, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, errors.New("Invalid or Expired Refresh Token.")
		}

		return 0, err
	}

	userId, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, errors.New("Invalid refresh token data.")
	}

	return userId, nil
}

func (s *RefreshTokenService) Delete(
	ctx context.Context,
	token string,
) error {
	if token == "" {
		return nil
	}

	tokenHash := security.HashRefreshToken(token)
	key := "refresh:" + tokenHash

	return s.redis.Del(ctx, key).Err()
}