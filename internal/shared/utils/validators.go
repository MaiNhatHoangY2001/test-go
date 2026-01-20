package utils

import (
	"regexp"

	errs "test-go/internal/shared/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ParseObjectID converts string to MongoDB ObjectID
func ParseObjectID(id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NilObjectID, errs.NewValidationError(errs.ErrIDRequired)
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, errs.NewValidationError(errs.ErrIDInvalid)
	}
	return objID, nil
}

// ValidateTitle validates a todo title
func ValidateTitle(title string) error {
	if title == "" {
		return errs.NewValidationError(errs.ErrTitleRequired)
	}
	if len(title) < 1 {
		return errs.NewValidationError(errs.ErrTitleTooShort)
	}
	if len(title) > 255 {
		return errs.NewValidationError(errs.ErrTitleTooLong)
	}
	return nil
}

// ValidateEmail validates an email address
func ValidateEmail(email string) error {
	if email == "" {
		return errs.NewValidationError(errs.ErrEmailRequired)
	}
	// Basic email regex validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errs.NewValidationError(errs.ErrEmailInvalid)
	}
	return nil
}

// ValidatePassword validates a password
func ValidatePassword(password string) error {
	if password == "" {
		return errs.NewValidationError(errs.ErrPasswordRequired)
	}
	if len(password) < 6 {
		return errs.NewValidationError(errs.ErrPasswordTooShort)
	}
	if len(password) > 128 {
		return errs.NewValidationError(errs.ErrPasswordTooLong)
	}
	return nil
}

// ValidateDescription validates todo description (optional but has max length)
func ValidateDescription(description string) error {
	if len(description) > 1000 {
		return errs.NewValidationError("description must not exceed 1000 characters")
	}
	return nil
}
