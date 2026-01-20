package utils

import (
	"regexp"

	errs "test-go/internal/shared/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ParseObjectID converts string to MongoDB ObjectID
func ParseObjectID(id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NilObjectID, errs.New(errs.ValidationError, errs.ErrIDRequired)
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, errs.New(errs.ValidationError, errs.ErrIDInvalid)
	}
	return objID, nil
}

// ValidateTitle validates a todo title
func ValidateTitle(title string) error {
	if title == "" {
		return errs.New(errs.ValidationError, errs.ErrTitleRequired)
	}
	if len(title) < 1 {
		return errs.New(errs.ValidationError, errs.ErrTitleTooShort)
	}
	if len(title) > 255 {
		return errs.New(errs.ValidationError, errs.ErrTitleTooLong)
	}
	return nil
}

// ValidateEmail validates an email address
func ValidateEmail(email string) error {
	if email == "" {
		return errs.New(errs.ValidationError, errs.ErrEmailRequired)
	}
	// Basic email regex validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errs.New(errs.ValidationError, errs.ErrEmailInvalid)
	}
	return nil
}

// ValidatePassword validates a password
func ValidatePassword(password string) error {
	if password == "" {
		return errs.New(errs.ValidationError, errs.ErrPasswordRequired)
	}
	if len(password) < 6 {
		return errs.New(errs.ValidationError, errs.ErrPasswordTooShort)
	}
	if len(password) > 128 {
		return errs.New(errs.ValidationError, errs.ErrPasswordTooLong)
	}
	return nil
}

// ValidateDescription validates todo description (optional but has max length)
func ValidateDescription(description string) error {
	if len(description) > 1000 {
		return errs.New(errs.ValidationError, "description must not exceed 1000 characters")
	}
	return nil
}
