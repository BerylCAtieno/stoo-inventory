package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

func ValidateOTP() bool {
	panic("ValidateOTP not implemented yet")
}
