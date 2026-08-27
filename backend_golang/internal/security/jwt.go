package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret        []byte
	accessExpires time.Duration
}

type AccessClaims struct {
	UserID 	uint64 	`json:"sub"`
	Role	string	`json:"role"`

	jwt.RegisteredClaims
}

func NewJWTService(
	secret string,
	accessExpires time.Duration,
) *JWTService {
	return &JWTService{
		secret: []byte(secret),
		accessExpires: accessExpires,
	}
}

func (s *JWTService) GenerateAccessToken (
	userID uint64,
	role string,
) (string, error) {
	now := time.Now()

	claims := AccessClaims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessExpires)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.secret)
}

func (s *JWTService) ParseAccessToken(
	tokenString string,
) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AccessClaims{},
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("Unexpected signing method.")
			}

			return s.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)

	if !ok {
		return nil, errors.New("Invalid claims")
	}

	return claims, nil
}

// func (s *JWTService) ValidateAccessToken(
// 	tokenString string,	
// ) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(
// 		tokenString,
// 		func(token *jwt.Token) (any, error) {
// 			if token.Method != jwt.SigningMethodHS256 {
// 				return nil, errors.New("Unexpected signing method.")
// 			}

// 			return s.secret, nil
// 		},
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, errors.New("Invalid token.")
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, errors.New("Invalid claims.")
// 	}

// 	return claims, nil
// }