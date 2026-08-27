package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func generateJWTSecret(size int) (string, error) {

	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to generate random bytes: %w", err)
	}
	
	secret := base64.URLEncoding.EncodeToString(bytes)
	
	return secret, nil
}

func generateJWTSecretHex(size int) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to generate random bytes: %w", err)
	}

	return fmt.Sprintf("%x", bytes), nil
}

func main() {
    secret, err := generateJWTSecret(32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("JWT_ACCESS_SECRET=%s\n", secret)
    
    longSecret, err := generateJWTSecret(64)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Long Secret: %s\n", longSecret)
    
    hexSecret, err := generateJWTSecretHex(32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Hex Secret: %s\n", hexSecret)
}