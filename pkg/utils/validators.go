package validators

import (
	"regexp"
	"strings"
)

const (
	minPasswordLength = 6
	maxPasswordLength = 50
	minTitleLength    = 1
	maxTitleLength    = 200
	maxDescLength     = 1000
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)
	return email != "" && emailRegex.MatchString(email) && len(email) <= 254
}

// ValidatePassword validates password strength
func ValidatePassword(password string) bool {
	if len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return false
	}
	return true
}

// ValidatePasswordStrength validates password contains mixed case and numbers
func ValidatePasswordStrength(password string) bool {
	if !ValidatePassword(password) {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

// ValidateName validates name field
func ValidateName(name string) bool {
	name = strings.TrimSpace(name)
	return len(name) > 0 && len(name) <= 100
}

// ValidateTodoTitle validates todo title
func ValidateTodoTitle(title string) bool {
	title = strings.TrimSpace(title)
	return len(title) >= minTitleLength && len(title) <= maxTitleLength
}

// ValidateTodoDescription validates todo description
func ValidateTodoDescription(description string) bool {
	return len(description) <= maxDescLength
}

// GetPasswordValidationError returns error message for invalid password
func GetPasswordValidationError(password string) string {
	if password == "" {
		return "password is required"
	}
	if len(password) < minPasswordLength {
		return "password must be at least " + string(rune(minPasswordLength)) + " characters"
	}
	if len(password) > maxPasswordLength {
		return "password must be at most " + string(rune(maxPasswordLength)) + " characters"
	}
	return ""
}

// GetEmailValidationError returns error message for invalid email
func GetEmailValidationError(email string) string {
	if email == "" {
		return "email is required"
	}
	if !ValidateEmail(email) {
		return "email format is invalid"
	}
	return ""
}
