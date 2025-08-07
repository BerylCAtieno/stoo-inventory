package utils

import (
	"fmt"

	"github.com/berylCAtieno/stoo-inventory/pkg/templates"
)

func SendPasswordResetEmail(email, otp string) {
	body, err := templates.LoadTemplate("reset_password.html", templates.EmailData{
		UserEmail: email,
		OTP:       otp,
	})
	if err != nil {
		// Handle error (log or return)
	}

	// Use `body` in your email-sending logic
	fmt.Println(body)
}
