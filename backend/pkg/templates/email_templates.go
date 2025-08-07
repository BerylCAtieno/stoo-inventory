package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

type EmailData struct {
	UserEmail string
	OTP       string
}

func LoadTemplate(templateName string, data EmailData) (string, error) {
	templatePath := filepath.Join("pkg", "templates", "templates", templateName)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return body.String(), nil
}
