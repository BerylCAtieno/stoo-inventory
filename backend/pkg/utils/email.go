package utils

import (
	"fmt"
	"log"

	"github.com/berylCAtieno/stoo-inventory/internal/config"
	"github.com/berylCAtieno/stoo-inventory/pkg/templates"
	"gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(email, otp string) error {
	body, err := templates.LoadTemplate("reset_password.html", templates.EmailData{
		UserEmail: email,
		OTP:       otp,
	})
	if err != nil {
		return fmt.Errorf("failed to load password reset email template: %w", err)
	}

	subject := "Password Reset Request"
	return sendEmail(email, subject, body)
}

func SendVerificationEmail(email, otp string) error {
	body, err := templates.LoadTemplate("verify_account.html", templates.EmailData{
		UserEmail: email,
		OTP:       otp,
	})
	if err != nil {
		return fmt.Errorf("failed to load verification email template: %w", err)
	}

	subject := "Verify Your Account"
	return sendEmail(email, subject, body)
}

func sendEmail(to, subject, body string) error {

	from := config.Config.MailAddress
	password := config.Config.SMTPPassword
	smtpHost := config.Config.SMTPHost
	smtpPort := config.Config.SMTPPort

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email to %s: %v", to, err)
		return err
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
