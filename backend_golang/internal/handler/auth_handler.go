package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/data/request"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/data/response"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/helper"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/validation"
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

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body.",
		)

		return
	}

	if err := validation.ValidateRegisterRequest(&req); err != nil {
		helper.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	err := h.authService.Register(
		r.Context(),
		req,
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
		map[string]string{
			"message": "Register Successful, Please Login...",
		},
	)
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req request.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)

		return
	}

	if err := validation.ValidateLoginRequest(&req); err != nil {
		helper.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	responseService, err := h.authService.Login(
		r.Context(),
		req,
	)
	if err != nil {
		helper.Error(
			w,
			http.StatusUnauthorized,
			err.Error(),
		)

		return
	}

	cookie := &http.Cookie{
		Name: "refresh_token",
		Value: responseService.RefreshToken,
		Path: "/api/v1/auth",
		HttpOnly: true,
		Secure: false,
		SameSite: http.SameSiteLaxMode,	
	}

	if req.RememberMe {
		cookie.MaxAge = int((7 * 24 * time.Hour).Seconds())
	}

	http.SetCookie(w, cookie)

	helper.JSON(
		w,
		http.StatusOK,
		response.LoginResponse{
			AccessToken: responseService.AccessToken,
			User: responseService.User,
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

	responseService, err := h.authService.Refresh(
		r.Context(),
		cookie.Value,
	)
	if err != nil {
		helper.Error(
			w,
			http.StatusUnauthorized,
			"Invalid or Expired refresh token.",
		)
		return
	}

	cookieRef := &http.Cookie{
		Name: "refresh_token",
		Value: responseService.RefreshToken,
		Path: "/api/v1/auth",
		HttpOnly: true,
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	}

	if responseService.RememberMe {
		cookieRef.MaxAge = int((7 * 24 * time.Hour).Seconds())
	}
	
	http.SetCookie(w, cookieRef)

	helper.JSON(
		w,
		http.StatusOK,
		response.LoginResponse{
			AccessToken: responseService.AccessToken,
			User: responseService.User,
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