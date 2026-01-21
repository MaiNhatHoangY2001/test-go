package constants

// Unique error codes for frontend i18n handling
// These are random unique codes, not HTTP status codes
const (
	CodeUnauthorized       = 10001
	CodeBadRequest         = 10002
	CodeForbidden          = 10003
	CodeNotFound           = 10004
	CodeInternalError      = 10005
	CodeValidationError    = 10006
	CodeConflict           = 10007
	CodeMissingAuthHeader  = 10008
	CodeInvalidAuthHeader  = 10009
	CodeInvalidToken       = 10010
	CodeInvalidTokenClaims = 10011
	CodeMissingUserID      = 10012
)

// HTTP status codes
const (
	StatusUnauthorized    = 401
	StatusBadRequest      = 400
	StatusForbidden       = 403
	StatusNotFound        = 404
	StatusInternalError   = 500
	StatusValidationError = 422
	StatusConflict        = 409
)

// HTTP error messages
const (
	MsgUnauthorized       = "UNAUTHORIZED"
	MsgBadRequest         = "BAD_REQUEST"
	MsgForbidden          = "FORBIDDEN"
	MsgNotFound           = "NOT_FOUND"
	MsgInternalError      = "INTERNAL_ERROR"
	MsgValidationError    = "VALIDATION_ERROR"
	MsgConflict           = "CONFLICT"
	MsgMissingAuthHeader  = "missing authorization header"
	MsgInvalidAuthHeader  = "invalid authorization header format"
	MsgInvalidToken       = "invalid or expired token"
	MsgInvalidTokenClaims = "invalid token claims"
	MsgMissingUserID      = "missing user_id in token"
)
