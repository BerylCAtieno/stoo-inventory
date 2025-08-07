package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/berylCAtieno/stoo-inventory/pkg/redisclient"
	"golang.org/x/net/context"
)

func GenerateOTP() (string, error) {
	max := big.NewInt(1000000)
	otp, err := rand.Int(rand.Reader, max)

	if err != nil {
		return "", err
	}

	otpString := fmt.Sprintf("%06d", otp.Int64())

	return otpString, nil
}

func StoreOtp(ctx context.Context, email string, otp string, ttl time.Duration) error {

	redisKey := fmt.Sprintf("otp:%s", email)

	err := redisclient.Client.Set(ctx, redisKey, otp, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to store OTP in Redis: %w", err)
	}

	return nil

}

func GetOtp(ctx context.Context, email string) (string, error) {
	redisKey := fmt.Sprintf("otp:%s", email)

	otp, err := redisclient.Client.Get(ctx, redisKey).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return "", fmt.Errorf("OTP not found or expired")
		}
		return "", fmt.Errorf("failed to retrieve OTP: %w", err)
	}

	return otp, nil
}

func DeleteOtp(ctx context.Context, email string) error {
	redisKey := fmt.Sprintf("otp:%s", email)
	return redisclient.Client.Del(ctx, redisKey).Err()
}

func VerifyOtp(email string, submittedOtp string) (bool, error) {
	ctx := context.Background()

	storedOtp, err := GetOtp(ctx, email)
	if err != nil {
		return false, err
	}

	if storedOtp == submittedOtp {
		err := DeleteOtp(ctx, email)
		if err != nil {
			log.Println("failed to delete OTP:", err)
		}
		return true, nil
	}

	return false, nil

}
