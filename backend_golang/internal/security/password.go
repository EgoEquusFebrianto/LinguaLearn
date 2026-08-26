package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PassowrdHasher struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
	saltLength  uint32
}

func NewPasswordHasher() *PassowrdHasher {
	return &PassowrdHasher{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		keyLength:   32,
		saltLength:  16,
	}
}

func (p *PassowrdHasher) Hash(password string) (string, error) {
	if password == "" {
		return "", errors.New("Password cannot be empty.")
	}

	salt := make([]byte, p.saltLength)

	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength,
	)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		p.memory,
		p.iterations,
		p.parallelism,
		encodedSalt,
		encodedHash,
	), nil
}

func (p *PassowrdHasher) Verify(
	passowrd string,
	encodedHash string,
) error {
	memory, iterations, parallelism, salt, expectedHash, err := parseHash(encodedHash)
	
	if err != nil {
		return err
	}

	actualHash := argon2.IDKey(
		[]byte(passowrd),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(expectedHash)),
	)

	if subtle.ConstantTimeCompare(actualHash, expectedHash) != 1 {
		return errors.New("Invalid Password.")
	}

	return nil
}

func parseHash(encodedHash string) (
	uint32,
	uint32,
	uint8,
	[]byte,
	[]byte,
	error,
) {
	parts := strings.Split(encodedHash, "$")

	if len(parts) != 6 {
		return 0, 0, 0, nil, nil, errors.New("Invalid passowrd hash format")
	}

	if parts[1] != "argon2id" {
		return 0, 0, 0, nil, nil, errors.New("Unsupported password algorithm")
	}

	var memory uint32
	var iterations uint32
	var parallelism uint8

	_, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&memory,
		&iterations,
		&parallelism,
	)

	if err != nil {
		return 0, 0, 0, nil, nil, errors.New("invalid argon2 parameters")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return 0, 0, 0, nil, nil, errors.New("invalid salt")
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return 0, 0, 0, nil, nil, errors.New("invalid hash")
	}

	return memory, iterations, parallelism, salt, expectedHash, nil
}