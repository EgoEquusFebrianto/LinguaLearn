package validation

import (
	"testing"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/delivery/http/request"
)

func TestValidation(t *testing.T) {
	inputRequest := &request.RegisterRequest{
		FullName: "John Doe",
		Email: "johndoe@example.com",
		Password: "password123",
		ConfirmPassword: "password123",
	}

	inputLogin := &request.LoginRequest{
		Email: "JohnDoe",
		Password: "pasword123",
		RememberMe: true,
	}

	err := ValidateRegisterRequest(inputRequest)
	if err != nil {
		t.Fatalf("expected no error in register, log.error: %v", err)
	}

	err = ValidateLoginRequest(inputLogin)
	if err != nil {
		t.Fatalf("expected no error in login, log.error: %v", err)
	}
}