package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/middleware"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/request"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/response"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/service"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils"
	"github.com/EgoEquusFebrianto/LinguaLearn/internal/utils/validation"
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
		utils.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body.",
		)

		return
	}

	if err := validation.ValidateRegisterRequest(&req); err != nil {
		utils.Error(
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
			utils.Error(
				w,
				http.StatusInternalServerError,
				"Internal server error",
			)
		default:
			utils.Error(
				w,
				http.StatusBadRequest,
				err.Error(),
			)
		}

		return
	}

	utils.JSON(
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
		utils.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)

		return
	}

	if err := validation.ValidateLoginRequest(&req); err != nil {
		utils.Error(
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
		utils.Error(
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

	utils.JSON(
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
		utils.Error(
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
		utils.Error(
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

	utils.JSON(
		w,
		http.StatusOK,
		response.LoginResponse{
			AccessToken: responseService.AccessToken,
			User: responseService.User,
		},
	)
}

func (h *AuthHandler) GetMe(
	w http.ResponseWriter, 
	r *http.Request,
) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		utils.Error(
			w,
			http.StatusUnauthorized,
			"User not found in context",
		)
		return
	}

	// role, _ := middleware.GetRole(r.Context())

	user, err := h.authService.GetMe(r.Context(), userID)
	if err != nil {
		utils.Error(
			w,
			http.StatusUnauthorized,
			"User not found.",
		)
	}

	utils.JSON(w, http.StatusOK, user)
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

	utils.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "Logged out successfully.",
		},
	)
}