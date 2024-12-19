package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func StoreVerificationCode(ctx context.Context, rdb *redis.Client, clientEmail string, code string, expiration time.Duration) error {
	key := fmt.Sprintf("verification_code:%s", clientEmail)
	err := rdb.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to store verification code: %v", err)
	}
	return nil
}

func ValidateVerificationCode(ctx context.Context, rdb *redis.Client, clientEmail string, inputCode string) error {
	key := fmt.Sprintf("verification_code:%s", clientEmail)

	storedCode, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("verification code not found or expired")
	} else if err != nil {
		return fmt.Errorf("failed to fetch verification code: %v", err)
	}

	if storedCode != inputCode {
		return fmt.Errorf("invalid verification code")
	}

	// err = rdb.Del(ctx, key).Err()
	// if err != nil {
	// 	return fmt.Errorf("failed to delete verification code: %v", err)
	// }

	return nil
}
