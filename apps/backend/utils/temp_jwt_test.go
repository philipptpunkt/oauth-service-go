package utils

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTemporaryJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	clientID := 123
	purpose := "email-verification"

	token, err := GenerateTemporaryJWT(clientID, purpose, time.Minute)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedEmail, parsedPurpose, err := ValidateTemporaryJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, clientID, parsedEmail)
	assert.Equal(t, purpose, parsedPurpose)
}

func TestValidateTemporaryJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	t.Run("ValidToken", func(t *testing.T) {
		token, _ := GenerateTemporaryJWT(123, "email-verification", time.Minute)
		clientID, purpose, err := ValidateTemporaryJWT(token)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if clientID != 123 || purpose != "email-verification" {
			t.Errorf("Expected email: 123 and purpose: email-verification")
		}
	})

	t.Run("InvalidToken", func(t *testing.T) {
		_, _, err := ValidateTemporaryJWT("invalid.token.string")
		if err == nil {
			t.Errorf("Expected an error for invalid token")
		}
	})
}
