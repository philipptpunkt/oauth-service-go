package utils

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTemporaryJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	email := "test@example.com"
	purpose := "email-verification"

	token, err := GenerateTemporaryJWT(email, purpose, time.Minute)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedEmail, parsedPurpose, err := ValidateTemporaryJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, email, parsedEmail)
	assert.Equal(t, purpose, parsedPurpose)
}

func TestValidateTemporaryJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	t.Run("ValidToken", func(t *testing.T) {
		token, _ := GenerateTemporaryJWT("test@example.com", "email-verification", time.Minute)
		email, purpose, err := ValidateTemporaryJWT(token)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if email != "test@example.com" || purpose != "email-verification" {
			t.Errorf("Expected email: test@example.com and purpose: email-verification")
		}
	})

	t.Run("InvalidToken", func(t *testing.T) {
		_, _, err := ValidateTemporaryJWT("invalid.token.string")
		if err == nil {
			t.Errorf("Expected an error for invalid token")
		}
	})
}
