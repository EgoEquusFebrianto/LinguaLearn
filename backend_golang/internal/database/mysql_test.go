package database

import (
	"os"
	"testing"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	"github.com/joho/godotenv"
)

func TestNewSQL(t *testing.T) {
	godotenv.Load("../../.env");

	cfg := &config.MySQLConfig{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := NewMySQL(cfg)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if db == nil {
		t.Error("Expected db not nil")
	}

	defer db.Close()
}