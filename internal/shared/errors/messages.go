package errors

// Common validation error messages
const (
	ErrTitleRequired    = "title is required"
	ErrTitleTooShort    = "title must be at least 1 character"
	ErrTitleTooLong     = "title must not exceed 255 characters"
	ErrEmailRequired    = "email is required"
	ErrEmailInvalid     = "email format is invalid"
	ErrPasswordRequired = "password is required"
	ErrPasswordTooShort = "password must be at least 6 characters"
	ErrPasswordTooLong  = "password must not exceed 128 characters"
	ErrIDInvalid        = "invalid id format"
	ErrIDRequired       = "id is required"
)

// Common application error messages
const (
	ErrTodoNotFound    = "todo not found"
	ErrUserNotFound    = "user not found"
	ErrUserExists      = "user already exists"
	ErrUnauthorized    = "unauthorized"
	ErrInvalidToken    = "invalid token"
	ErrTokenExpired    = "token expired"
	ErrInvalidPassword = "invalid password"
	ErrDatabaseError   = "database error occurred"
	ErrInternalError   = "internal server error"
)
