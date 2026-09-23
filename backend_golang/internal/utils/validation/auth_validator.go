package validation

import (
	"errors"
	"strings"
	"unicode"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/request"
)

func ValidateRegisterRequest(req *request.RegisterRequest) error {
	fullName := strings.TrimSpace(req.FullName)
	
	if fullName == "" {
		return errors.New("Full name is required.")
	}

	if isAlphabetic(fullName) {
		return errors.New("Full name must contain alphabetic characters or spaces only.")
	}

	if req.Email == "" {
		return errors.New("Email is required.")
	}

	if hasWhiteSpace(req.Email) {
		return errors.New("Email should not contain spaces.")
	}

	if req.Password == "" {
		return errors.New("Password is required.")
	}

	if len([]rune(req.Password)) < 8 {
		return errors.New("Password should be at least 8 characters.")
	}

	if req.ConfirmPassword == "" {
		return errors.New("Password confirmation required.")
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("Passwords doesn't match.")
	}

	return nil
}

func ValidateLoginRequest(req *request.LoginRequest) error {
	if req.Email == "" {
		return errors.New("email is required.")
	}

	if hasWhiteSpace(req.Email) {
		return errors.New("email must not contain spaces.")
	}

	if req.Password == "" {
		return errors.New("password is required.")
	}

	return nil
}

func isAlphabetic(value string) bool {
	for _, r := range value {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == ' ') {
			return true
		}
	}

	return false
}

func hasWhiteSpace(value string) bool {
	for _, r := range value {
		if unicode.IsSpace(r) {
			return true
		}
	}

	return false
}