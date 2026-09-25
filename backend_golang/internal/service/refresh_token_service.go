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
	RememberMe bool,
) (string, error) {
	token, tokenHash, err := security.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	key := "refresh:" + tokenHash
	
	_, err = s.redis.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(
			ctx,
			key,
			"user_id", UserID,
			"remember_me", RememberMe,
		)

		pipe.Expire(
			ctx,
			key,
			refreshTokenTTL,
		)

		return nil
	})
	
	// err = s.redis.Set(
	// 	ctx,
	// 	key,
	// 	strconv.FormatUint(UserID, 10),
	// 	refreshTokenTTL,
	// ).Err()

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *RefreshTokenService) GetUserData(
	ctx context.Context,
	token string,
) (uint64, bool, error) {
	if token == "" {
		return 0, false, errors.New("Refresh token is required.")
	}

	tokenHash := security.HashRefreshToken(token)
	key := "refresh:" + tokenHash

	// value, err := s.redis.Get(ctx, key).Result()
	// if err != nil {
	// 	if errors.Is(err, redis.Nil) {
	// 		return 0, false, errors.New("Invalid or Expired Refresh Token.")
	// 	}

	// 	return 0, false, err
	// }

	value, err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return 0, false, err
	}
	
	if len(value) == 0 {
		return 0, false, errors.New("Invalid or Expired Refresh Token.")
	}

	userId, err := strconv.ParseUint(value["user_id"], 10, 64)
	if err != nil {
		return 0, false, errors.New("Invalid refresh token data.")
	}

	rememberMe := value["remember_me"] == "1"

	return userId, rememberMe, nil
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