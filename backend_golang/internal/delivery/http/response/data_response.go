package response

import "github.com/EgoEquusFebrianto/LinguaLearn/internal/domain"

type LoginServiceResponse struct {
	AccessToken  string
	RefreshToken string
	User         UserProfile
}

type RefreshServiceResponse struct {
	AccessToken  string
	RefreshToken string
	User         UserProfile
	RememberMe   bool
}

type UserProfile struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	AccessToken string      `json:"access_token"`
	User        interface{} `json:"user"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type BankWordSearchResponse struct {
	Data       []domain.BankWord `json:"data"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	Total      int64             `json:"total"`
	TotalPages int               `json:"total_pages"`
}
