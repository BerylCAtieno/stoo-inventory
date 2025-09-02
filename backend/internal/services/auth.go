package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/berylCAtieno/stoo-inventory/internal/config"
	"github.com/berylCAtieno/stoo-inventory/internal/models"
	"github.com/berylCAtieno/stoo-inventory/internal/repositories"
	"github.com/berylCAtieno/stoo-inventory/pkg/utils"
)

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Register user
func (s *AuthService) Register(firstName, lastName, email, password string) error {
	// Check if user exists
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}
	if !errors.Is(err, repositories.ErrNotFound) {
		return fmt.Errorf("db error: %w", err)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: hashedPassword,
		IsActive:     false,
	}

	// Save user
	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Generate OTP
	otp, err := utils.GenerateOtp()
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	// Send verification email
	if err := utils.SendVerificationEmail(email, otp); err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	// Store OTP in Redis
	ctx := context.Background()
	if err := utils.StoreOtp(ctx, email, otp, config.Config.OtpExpiry); err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	return nil
}

// Verify OTP
func (s *AuthService) VerifyOTP(email, otp string) error {
	ok, err := utils.VerifyOtp(email, otp)
	if err != nil {
		return fmt.Errorf("failed to verify otp: %w", err)
	}
	if !ok {
		return errors.New("invalid or expired OTP")
	}

	return s.repo.UpdateStatus(email, map[string]interface{}{
		"is_active":  true,
		"updated_at": time.Now(),
	})
}

// Login user
func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return "", "", errors.New("invalid credentials")
		}
		return "", "", fmt.Errorf("db error: %w", err)
	}

	if !user.IsActive {
		return "", "", errors.New("account not activated")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		time.Sleep(500 * time.Millisecond) // slow down brute-force
		return "", "", errors.New("invalid credentials")
	}

	// Generate JWT pair
	accessToken, refreshToken, err := utils.GenerateJWTPair(
		user.ID,
		config.Config.JwtSecret,
		config.Config.JwtExpirationHours,
		config.Config.JwtRefreshHours,
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate jwt: %w", err)
	}

	// Update last login
	if err := s.repo.UpdateStatus(email, map[string]interface{}{
		"last_login_at": time.Now(),
		"updated_at":    time.Now(),
	}); err != nil {
		return "", "", fmt.Errorf("failed to update last login: %w", err)
	}

	return accessToken, refreshToken, nil
}

// ForgotPassword - send OTP for reset
func (s *AuthService) ForgotPassword(email string) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return errors.New("no account found with that email")
		}
		return fmt.Errorf("db error: %w", err)
	}

	otp, err := utils.GenerateOtp()
	if err != nil {
		return fmt.Errorf("failed to generate otp: %w", err)
	}

	if err := utils.SendPasswordResetEmail(user.Email, otp); err != nil {
		return fmt.Errorf("failed to send reset email: %w", err)
	}

	ctx := context.Background()
	if err := utils.StoreOtp(ctx, email, otp, config.Config.ResetOtpExpiry); err != nil {
		return fmt.Errorf("failed to store reset otp: %w", err)
	}

	return nil
}

// ResetPassword - verify OTP and update password
func (s *AuthService) ResetPassword(email, otp, newPassword string) error {
	ok, err := utils.VerifyOtp(email, otp)
	if err != nil {
		return fmt.Errorf("failed to verify otp: %w", err)
	}
	if !ok {
		return errors.New("invalid or expired OTP")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return s.repo.UpdateStatus(email, map[string]interface{}{
		"password_hash": hashedPassword,
		"updated_at":    time.Now(),
	})
}
