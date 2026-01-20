package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode represents different error types in the application
type ErrorCode string

const (
	// Client errors
	ValidationError   ErrorCode = "VALIDATION_ERROR"
	NotFoundError     ErrorCode = "NOT_FOUND"
	ConflictError     ErrorCode = "CONFLICT"
	UnauthorizedError ErrorCode = "UNAUTHORIZED"
	ForbiddenError    ErrorCode = "FORBIDDEN"
	BadRequestError   ErrorCode = "BAD_REQUEST"

	// Server errors
	InternalError ErrorCode = "INTERNAL_ERROR"
	DatabaseError ErrorCode = "DATABASE_ERROR"
	ExternalError ErrorCode = "EXTERNAL_ERROR"
)

// AppError represents a standardized application error
type AppError struct {
	Code    ErrorCode
	Message string
	Details string
	err     error
}

// New creates a new AppError
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// NewWithDetails creates a new AppError with additional details
func NewWithDetails(code ErrorCode, message, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Wrap wraps an existing error with an AppError
func Wrap(err error, code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		err:     err,
	}
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.err
}

// HTTPStatusCode returns the appropriate HTTP status code for this error
func (e *AppError) HTTPStatusCode() int {
	switch e.Code {
	case ValidationError, BadRequestError:
		return http.StatusBadRequest
	case UnauthorizedError:
		return http.StatusUnauthorized
	case ForbiddenError:
		return http.StatusForbidden
	case NotFoundError:
		return http.StatusNotFound
	case ConflictError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == ValidationError
	}
	return false
}

// IsNotFoundError checks if the error is a not found error
func IsNotFoundError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == NotFoundError
	}
	return false
}

// IsConflictError checks if the error is a conflict error
func IsConflictError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == ConflictError
	}
	return false
}

// IsDatabaseError checks if the error is a database error
func IsDatabaseError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == DatabaseError
	}
	return false
}
