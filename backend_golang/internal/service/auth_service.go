package service

import (
	"context"
	"errors"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/request"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/response"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/domain"
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
) error {
	_, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err == nil {
		return errors.New("Email already registered.")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	passwordHash, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return err
	}

	user := &domain.User{
		RoleId: 1,
		FullName: req.FullName,
		Email: req.Email,
		PasswordHash: passwordHash,
	}
	if err := s.userRepository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(
	ctx context.Context,
	req request.LoginRequest,
) (*response.LoginServiceResponse, error ) {
	user, err := s.userRepository.FindByEmail(ctx, req.Email)
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
	userID, rememberMe, err := s.refreshTokenService.GetUserData(
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
		rememberMe,
	)

	if err!= nil {
		return nil, err
	}

	return &response.RefreshServiceResponse{
		AccessToken: accessToken,
		RefreshToken: newRefreshToken,
		User: response.UserProfile{
			ID: userID,
			FullName: user.FullName,
			Email: user.Email,
			Role: user.Role.Name,
		},
		RememberMe: rememberMe,
	}, nil
}

func (s *AuthService) GetMe(
	ctx context.Context,
	userID uint64,
) (*response.UserProfile, error) {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userProfile := &response.UserProfile{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		Role: user.Role.Name,
	}

	return userProfile, nil
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