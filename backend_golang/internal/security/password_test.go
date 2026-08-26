package security

import "testing"

func TestPasswordHasher(t *testing.T) {
	hasher := NewPasswordHasher()

	password := "my-secret-password"

	hash, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if hash == password {
		t.Fatal("Password must be not stored as plaintext")
	}

	if err := hasher.Verify(password, hash); err != nil {
		t.Fatalf("Verify() failed for context password: %v", err)
	}

	if err := hasher.Verify("wrong-password", hash); err == nil {
		t.Fatal("Verify() should fail for incorrect password")
	}
}

func TestPasswordHasherDifferentSalt(t *testing.T) {
	hasher := NewPasswordHasher()

	password := "my-secret-password"

	hash1, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	hash2, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if hash1 == hash2 {
		t.Fatal("same password should produce different hashes because of random salt")
	}
}