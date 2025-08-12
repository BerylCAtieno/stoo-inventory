package services

import (
	"context"
	"errors"
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
	if err := s.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// Create user
	user := models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: hashedPassword,
		IsActive:     false, // inactive until OTP verified
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}

	// Generate OTP
	otp, err := utils.GenerateOtp()
	if err != nil {
		return err
	}

	// Send verification email
	if err := utils.SendVerificationEmail(email, otp); err != nil {
		return err
	}

	// Store OTP in Redis
	ctx := context.Background()
	if err := utils.StoreOtp(ctx, email, otp, 5*time.Minute); err != nil {
		return err
	}

	return nil
}

// Verify OTP
func (s *AuthService) VerifyOTP(email, otp string) error {
	ok, err := utils.VerifyOtp(email, otp)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid or expired OTP")
	}

	// Activate user
	return s.DB.Model(&models.User{}).Where("email = ?", email).Update("is_active", true).Error
}

// Login user
func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", errors.New("account not activated")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, config.Config.JwtSecret, config.Config.JwtExpirationHours)
	if err != nil {
		return "", err
	}

	// Update last login time
	now := time.Now()
	s.DB.Model(&user).Update("last_login_at", now)

	return token, nil
}

// ForgotPassword - send OTP to reset password
func (s *AuthService) ForgotPassword(email string) error {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("no account found with that email")
	}

	// Generate OTP
	otp, err := utils.GenerateOtp()
	if err != nil {
		return err
	}

	// Send password reset email
	if err := utils.SendPasswordResetEmail(email, otp); err != nil {
		return err
	}

	// Store OTP in Redis
	ctx := context.Background()
	if err := utils.StoreOtp(ctx, email, otp, 10*time.Minute); err != nil {
		return err
	}

	return nil
}

// ResetPassword - verify OTP and update password
func (s *AuthService) ResetPassword(email, otp, newPassword string) error {
	ok, err := utils.VerifyOtp(email, otp)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid or expired OTP")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password in DB
	return s.DB.Model(&models.User{}).
		Where("email = ?", email).
		Update("password_hash", hashedPassword).Error
}
