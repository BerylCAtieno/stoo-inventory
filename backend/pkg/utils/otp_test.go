package utils

import (
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	otp, err := GenerateOTP()

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(otp) != 6 {
		t.Errorf("Expected OTP length to be 6, got: %d (OTP: %s)", len(otp), otp)
	}

	for _, ch := range otp {
		if ch < '0' || ch > '9' {
			t.Errorf("Expected OTP to contain only digits, got: %s", otp)
			break
		}
	}
}
