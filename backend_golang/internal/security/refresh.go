package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateRefreshToken() (string, string, error) {
	bytes := make([]byte, 64)

	if _, err := rand.Read(bytes); err != nil {
		return "", "", err
	}

	token := base64.RawURLEncoding.EncodeToString(bytes)

	hash := sha256.Sum256([]byte(token))
	tokenHash := base64.RawURLEncoding.EncodeToString(hash[:])

	return token, tokenHash, nil
}

func HashRefreshToken(token string) string {
	hash := sha256.Sum256([]byte(token))

	return base64.RawURLEncoding.EncodeToString(hash[:])
}