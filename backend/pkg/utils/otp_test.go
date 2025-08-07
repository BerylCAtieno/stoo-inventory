package utils

import (
	"testing"
	"time"

	"golang.org/x/net/context"
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

func TestOtpFlow(t *testing.T) {
	ctx := context.Background()
	email := "test@example.com"

	// Step 1: Generate OTP
	otp, err := GenerateOTP()
	if err != nil {
		t.Fatalf("GenerateOTP failed: %v", err)
	}

	if len(otp) != 6 {
		t.Errorf("Expected OTP length of 6, got %d", len(otp))
	}

	// Step 2: Store OTP
	ttl := 10 * time.Second
	err = StoreOtp(ctx, email, otp, ttl)
	if err != nil {
		t.Fatalf("StoreOtp failed: %v", err)
	}

	// Step 3: Get OTP
	storedOtp, err := GetOtp(ctx, email)
	if err != nil {
		t.Fatalf("GetOtp failed: %v", err)
	}

	if storedOtp != otp {
		t.Errorf("Expected stored OTP %s, got %s", otp, storedOtp)
	}

	// Step 4: Verify OTP - should be true
	isValid, err := VerifyOtp(email, otp)
	if err != nil {
		t.Fatalf("VerifyOtp failed: %v", err)
	}

	if !isValid {
		t.Error("VerifyOtp returned false, expected true")
	}

	// Step 5: Verify again - should fail because OTP is deleted
	isValid, err = VerifyOtp(email, otp)
	if err == nil {
		t.Error("Expected error after OTP was deleted, got none")
	}

	if isValid {
		t.Error("VerifyOtp returned true on second attempt, expected false")
	}
}

func TestDeleteOtp(t *testing.T) {
	ctx := context.Background()
	email := "delete@test.com"
	otp := "123456"

	// Store a dummy OTP
	err := StoreOtp(ctx, email, otp, 30*time.Second)
	if err != nil {
		t.Fatalf("StoreOtp failed: %v", err)
	}

	// Delete the OTP
	err = DeleteOtp(ctx, email)
	if err != nil {
		t.Fatalf("DeleteOtp failed: %v", err)
	}

	// Try to get it
	_, err = GetOtp(ctx, email)
	if err == nil {
		t.Error("Expected error after OTP was deleted, got none")
	}
}

//TODO: Test OTP flow properly
