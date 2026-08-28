package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type registerRequest struct {
	FullName string	`json:"full_name"`
	Email    string	`json:"email"`
	Password string	`json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
    AccessToken string      `json:"access_token"`
    User        interface{} `json:"user"`
}

type userResponse struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body.",
		)
	}

	user, err := h.authService.Register(
		r.Context(),
		service.RegisterRequest{
			FullName: req.FullName,
			Email: req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helper.Error(
				w,
				http.StatusInternalServerError,
				"Internal server error",
			)
		default:
			helper.Error(
				w,
				http.StatusBadRequest,
				err.Error(),
			)
		}

		return
	}

	helper.JSON(
		w,
		http.StatusCreated,
		userResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
		},
	)
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)

		return
	}

	response, err := h.authService.Login(
		r.Context(),
		service.LoginRequest{
			Email: req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		helper.Error(
			w,
			http.StatusUnauthorized,
			err.Error(),
		)

		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name: "refresh_token",
			Value: response.RefreshToken,
			Path: "/api/v1/auth",
			HttpOnly: true,
			Secure: false,
			SameSite: http.SameSiteLaxMode,
			MaxAge: int((7 * 24 * time.Hour).Seconds()),
		},
	)

	helper.JSON(
		w,
		http.StatusOK,
		loginResponse{
			AccessToken: response.AccessToken,
			User: response.User,
		},
	)
}

func (h *AuthHandler) Refresh (
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		helper.Error(
			w,
			http.StatusUnauthorized,
			"Refresh token required.",
		)
		return
	}

	response, err := h.authService.Refresh(
		r.Context(),
		cookie.Value,
	)
	if err != nil {
		helper.Error(
			w,
			http.StatusUnauthorized,
			"Invalid or Expired refresh token.",
		)
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name: "refresh_token",
			Value: response.RefreshToken,
			Path: "/api/v1/auth",
			HttpOnly: true,
			Secure: false,
			SameSite: http.SameSiteLaxMode,
			MaxAge: int((7 * 24 * time.Hour).Seconds()),
		},
	)

	helper.JSON(
		w,
		http.StatusOK,
		loginResponse{
			AccessToken: response.AccessToken,
			User: response.User,
		},
	)
}

func (h *AuthHandler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, err := r.Cookie("refresh_token")
	if err == nil {
		_ = h.authService.Logout(
			r.Context(),
			cookie.Value,
		)
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name: "refresh_token",
			Value: "",
			Path: "/api/v1/auth",
			HttpOnly: true,
			Secure: false,
			SameSite: http.SameSiteLaxMode,
			MaxAge: -1,
		},
	)

	helper.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "Logged out successfully.",
		},
	)
}