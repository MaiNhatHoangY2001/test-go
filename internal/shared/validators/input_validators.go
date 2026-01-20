package validators

import (
	"regexp"
	"strings"
	"unicode"
)

// ValidationError represents validation errors with field details
type ValidationError struct {
	Field   string
	Message string
}

// EmailValidator validates email format
func EmailValidator(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email) && len(email) <= 254
}

// PasswordValidator validates password strength
// Requirements: min 8 chars, 1 uppercase, 1 lowercase, 1 digit, 1 special char
func PasswordValidator(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}

	if len(password) > 128 {
		return false, "Password must not exceed 128 characters"
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return false, "Password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "Password must contain at least one lowercase letter"
	}
	if !hasDigit {
		return false, "Password must contain at least one digit"
	}
	if !hasSpecial {
		return false, "Password must contain at least one special character"
	}

	return true, ""
}

// TodoTitleValidator validates todo title
func TodoTitleValidator(title string) (bool, string) {
	title = strings.TrimSpace(title)
	
	if len(title) == 0 {
		return false, "Title is required"
	}

	if len(title) < 3 {
		return false, "Title must be at least 3 characters long"
	}

	if len(title) > 255 {
		return false, "Title must not exceed 255 characters"
	}

	return true, ""
}

// TodoDescriptionValidator validates todo description
func TodoDescriptionValidator(description string) (bool, string) {
	if len(description) > 2000 {
		return false, "Description must not exceed 2000 characters"
	}
	return true, ""
}

// StringLengthValidator validates string length
func StringLengthValidator(value string, minLen, maxLen int) (bool, string) {
	value = strings.TrimSpace(value)
	
	if len(value) < minLen {
		return false, "Value must be at least " + string(rune(minLen)) + " characters"
	}

	if len(value) > maxLen {
		return false, "Value must not exceed " + string(rune(maxLen)) + " characters"
	}

	return true, ""
}

// PaginationValidator validates pagination parameters
func PaginationValidator(page, limit int) (bool, string) {
	if page < 1 {
		return false, "Page must be at least 1"
	}

	if limit < 1 {
		return false, "Limit must be at least 1"
	}

	if limit > 100 {
		return false, "Limit must not exceed 100"
	}

	return true, ""
}
