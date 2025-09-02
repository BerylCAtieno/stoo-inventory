package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/berylCAtieno/stoo-inventory/internal/config"
	"github.com/berylCAtieno/stoo-inventory/internal/models"
	"github.com/berylCAtieno/stoo-inventory/pkg/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

// Register user
func (s *AuthService) Register(firstName, lastName, email, password string) error {
	// Check if user exists
	var existing models.User
	err := s.DB.Where("email = ?", email).First(&existing).Error
	if err == nil {
		return errors.New("user already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("db error: %w", err)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: hashedPassword,
		IsActive:     false,
	}

	// Run transaction
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		otp, err := utils.GenerateOtp()
		if err != nil {
			return fmt.Errorf("failed to generate OTP: %w", err)
		}

		if err := utils.SendVerificationEmail(email, otp); err != nil {
			return fmt.Errorf("failed to send verification email: %w", err)
		}

		ctx := context.Background()
		if err := utils.StoreOtp(ctx, email, otp, config.Config.OtpExpiry); err != nil {
			return fmt.Errorf("failed to store OTP: %w", err)
		}

		return nil
	})
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

	return s.DB.Model(&models.User{}).
		Where("email = ?", email).
		Update("is_active", true).Error
}

// Login user
func (s *AuthService) Login(email, password string) (string, string, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", errors.New("invalid credentials")
		}
		return "", "", fmt.Errorf("db error: %w", err)
	}

	if !user.IsActive {
		return "", "", errors.New("account not activated")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		// slow down brute-force
		time.Sleep(500 * time.Millisecond)
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
	now := time.Now()
	if err := s.DB.Model(&user).Update("last_login_at", now).Error; err != nil {
		return "", "", fmt.Errorf("failed to update last login: %w", err)
	}

	return accessToken, refreshToken, nil
}

// ForgotPassword - send OTP for reset
func (s *AuthService) ForgotPassword(email string) error {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no account found with that email")
		}
		return fmt.Errorf("db error: %w", err)
	}

	otp, err := utils.GenerateOtp()
	if err != nil {
		return fmt.Errorf("failed to generate otp: %w", err)
	}

	if err := utils.SendPasswordResetEmail(email, otp); err != nil {
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

	return s.DB.Model(&models.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"password_hash": hashedPassword,
			"updated_at":    time.Now(),
		}).Error
}
