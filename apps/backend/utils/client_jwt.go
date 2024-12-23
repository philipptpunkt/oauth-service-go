package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateClientJWT(userID int, email string, purpose string, duration time.Duration) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{
		"userID":  userID,
		"email":   email,
		"purpose": purpose,
		"exp":     time.Now().Add(duration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ValidateClientJWT(tokenString string) (int, string, string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return 0, "", "", fmt.Errorf("JWT_SECRET is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, "", "", fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", "", fmt.Errorf("invalid token claims")
	}

	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, "", "", fmt.Errorf("missing or invalid userID in token")
	}
	userID := int(userIDFloat)

	email, ok := claims["email"].(string)
	if !ok {
		return 0, "", "", fmt.Errorf("missing or invalid email in token")
	}

	purpose, ok := claims["purpose"].(string)
	if !ok {
		return 0, "", "", fmt.Errorf("missing or invalid purpose in token")
	}

	return userID, email, purpose, nil
}
