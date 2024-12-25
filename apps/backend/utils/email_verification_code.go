package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func StoreVerificationCode(ctx context.Context, rdb *redis.Client, clientID int, code string, expiration time.Duration) error {
	key := fmt.Sprintf("verification_code:%d", clientID) // Corrected to %d for integers
	err := rdb.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to store verification code: %v", err)
	}
	return nil
}

func ValidateVerificationCode(ctx context.Context, rdb *redis.Client, clientID int, inputCode string) error {
	key := fmt.Sprintf("verification_code:%d", clientID) // Corrected to %d for integers

	storedCode, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("verification code not found or expired")
	} else if err != nil {
		return fmt.Errorf("failed to fetch verification code: %v", err)
	}

	if storedCode != inputCode {
		return fmt.Errorf("invalid verification code")
	}

	// Optionally delete the code after validation
	// err = rdb.Del(ctx, key).Err()
	// if err != nil {
	// 	return fmt.Errorf("failed to delete verification code: %v", err)
	// }

	return nil
}
