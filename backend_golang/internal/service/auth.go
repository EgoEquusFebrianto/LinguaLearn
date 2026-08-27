package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/models"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	FullName string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	userRepository 	repository.UserRepository
	passwordHasher 	*security.PasswordHasher
	jwtService 		*security.JWTService
	redis			*redis.Client
}

func NewauthService(
	userRepository repository.UserRepository,
	passwordHasher *security.PasswordHasher,
	jwtService *security.JWTService,
	redis *redis.Client,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		jwtService: jwtService,
		redis: redis,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	req RegisterRequest,
) (*models.User, error) {
	fullName := strings.TrimSpace((req.FullName))
	email := strings.ToLower(strings.TrimSpace(req.Email))

	if fullName == "" {
		return nil, errors.New("Fullname is required.")
	}

	if email == "" {
		return nil, errors.New("Email is required.")
	}

	if req.Password == "" {
		return nil, errors.New("Password is required.")
	}

	_, err := s.userRepository.FindByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("Email already registered.")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	passwordHash, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		RoleId: 1,
		FullName: fullName,
		Email: email,
		PasswordHash: passwordHash,
	}
	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	req LoginRequest,
) (*LoginResponse, error ) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	if email == "" {
		return nil, errors.New("Email is required.")
	}

	if req.Password == "" {
		return nil, errors.New("Passowrd is required.")
	}

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Invalid email or password.")
		}

		return nil, err
	}

	if err := s.passwordHasher.Verify(
		req.Password,
		user.PasswordHash,
	); err != nil {
		return nil, errors.New("Invalid email or password.")
	}

	accessToken, err := s.jwtService.GenerateAccessToken(
		user.ID,
		user.Role.Name,
	)

	if err != nil {
		return nil, err
	}

	refreshToken, refreshTokenHash, err := security.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	key := "refresh:" + refreshTokenHash

	err = s.redis.Set(
		ctx,
		key,
		user.ID,
		7 * 24 * time.Hour,
	).Err()

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}