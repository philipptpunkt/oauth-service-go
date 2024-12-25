package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	RoleAdmin      Role = "admin"
	RoleMaintainer Role = "maintainer"
	RoleViewer     Role = "viewer"
	RoleSupport    Role = "support"
	RoleOwner      Role = "owner"
)

func GenerateClientJWT(clientID int, role Role, duration time.Duration) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{
		"clientID": clientID,
		"role":     role,
		"exp":      time.Now().Add(duration).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ValidateClientJWT(tokenString string) (int, Role, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return 0, "", fmt.Errorf("JWT_SECRET is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, "", fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	clientIDFloat, ok := claims["clientID"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("missing or invalid clientID in token")
	}
	clientID := int(clientIDFloat)

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", fmt.Errorf("missing or invalid role in token")
	}

	switch Role(role) {
	case RoleAdmin, RoleMaintainer, RoleViewer, RoleSupport, RoleOwner:
		return clientID, Role(role), nil
	default:
		return 0, "", fmt.Errorf("unauthorized role in token")
	}
}
