package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTCreateAndValidate(t *testing.T) {
	secret := "testsecret"
	userID := uuid.New()

	token, err := MakeJWT(userID, secret, time.Minute)
	if err != nil {
		t.Fatalf("failed to make jwt: %v", err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("failed to validate jwt: %v", err)
	}

	if gotID != userID {
		t.Errorf("expected %v, got %v", userID, gotID)
	}
}

func TestJWTExpired(t *testing.T) {
	secret := "testsecret"
	userID := uuid.New()

	token, err := MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("failed to make jwt: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Errorf("expected error for expired token, got nil")
	}
}

func TestJWTWrongSecret(t *testing.T) {
	secret := "testsecret"
	userID := uuid.New()

	token, err := MakeJWT(userID, secret, time.Minute)
	if err != nil {
		t.Fatalf("failed to make jwt: %v", err)
	}

	_, err = ValidateJWT(token, "wrongsecret")
	if err == nil {
		t.Errorf("expected error for wrong secret, got nil")
	}
}
