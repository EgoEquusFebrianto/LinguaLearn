package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func LoadEnv(t *testing.T) {
	t.Helper()

	dir, _ := os.Getwd()
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				t.Fatal(err)
			}
			return
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal(".env tidak ditemukan")
		}
		dir = parent
	}
}