package service

import (
	"context"
	"errors"
	"strings"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/data/request"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/data/response"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/models"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/repository"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository 		repository.UserRepository
	passwordHasher 		*security.PasswordHasher
	jwtService 			*security.JWTService
	refreshTokenService	*RefreshTokenService
}

func NewAuthService(
	userRepository repository.UserRepository,
	passwordHasher *security.PasswordHasher,
	jwtService *security.JWTService,
	refreshTokenService *RefreshTokenService,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		jwtService: jwtService,
		refreshTokenService: refreshTokenService,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	req request.RegisterRequest,
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
	req request.LoginRequest,
) (*response.LoginServiceResponse, error ) {
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

	refreshToken, err := s.refreshTokenService.Create(
		ctx,
		user.ID,
		req.RememberMe,
	)
	if err != nil {
		return nil, err
	}

	return &response.LoginServiceResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
        User: response.UserProfile{
            ID:       user.ID,
            FullName: user.FullName,
            Email:    user.Email,
            Role:     user.Role.Name,
        },
	}, nil
}

func (s *AuthService) Refresh(
	ctx context.Context,
	refreshToken string,
) (*response.RefreshServiceResponse, error) {
	userID, remember_me, err := s.refreshTokenService.GetUserData(
		ctx,
		refreshToken,
	)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtService.GenerateAccessToken(
		user.ID,
		user.Role.Name,
	)
	if err != nil {
		return nil, err
	}

	// Refresh Token Rotation
	err = s.refreshTokenService.Delete(
		ctx,
		refreshToken,
	)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.refreshTokenService.Create(
		ctx,
		user.ID,
		remember_me,
	)

	return &response.RefreshServiceResponse{
		AccessToken: accessToken,
		RefreshToken: newRefreshToken,
		User: response.UserProfile{
			ID: userID,
			FullName: user.FullName,
			Email: user.Email,
			Role: user.Role.Name,
		},
		RememberMe: remember_me,
	}, nil
}

func (s *AuthService) Logout(
	ctx context.Context,
	refreshToken string,
) error {
	return s.refreshTokenService.Delete(
		ctx,
		refreshToken,
	)
}