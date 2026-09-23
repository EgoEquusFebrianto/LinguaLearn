package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/security"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils"
)

// CONSTANTS
type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

// ERRORS 
var (
	ErrMissingAuthHeader = errors.New("authorization header required")
	ErrInvalidAuthHeader = errors.New("invalid authorization header format")
)

// MIDDLEWARE 

// Auth - Authentication middleware
func Auth(jwtService *security.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(
			w http.ResponseWriter, 
			r *http.Request) {
				token, err := extractToken(r)
				if err != nil {
					utils.Error(w, http.StatusUnauthorized, err.Error())
					return
				}

				claims, err := jwtService.ParseAccessToken(token)
				if err != nil {
					utils.Error(w, http.StatusUnauthorized, "Invalid or expired access token")
					return
				}

				ctx := setAuthContext(r.Context(), claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}

// utilsS 

// extractToken - Extract Bearer token from Authorization header
func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrMissingAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

// setAuthContext - Add user claims to context
func setAuthContext(ctx context.Context, claims *security.AccessClaims) context.Context {
	ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
	ctx = context.WithValue(ctx, RoleKey, claims.Role)
	return ctx
}

// GetUserID - Get user ID from context
func GetUserID(ctx context.Context) (uint64, bool) {
	userID, ok := ctx.Value(UserIDKey).(uint64)
	return userID, ok
}

// GetRole - Get user role from context
func GetRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(RoleKey).(string)
	return role, ok
}