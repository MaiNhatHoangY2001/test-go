package constants

// HTTP error codes
const (
	CodeUnauthorized       = 401
	CodeBadRequest         = 400
	CodeForbidden          = 403
	CodeNotFound           = 404
	CodeInternalError      = 500
	CodeValidationError    = 422
	CodeConflict           = 409
)

// HTTP error messages
const (
	MsgUnauthorized           = "UNAUTHORIZED"
	MsgBadRequest             = "BAD_REQUEST"
	MsgForbidden              = "FORBIDDEN"
	MsgNotFound               = "NOT_FOUND"
	MsgInternalError          = "INTERNAL_ERROR"
	MsgValidationError        = "VALIDATION_ERROR"
	MsgConflict               = "CONFLICT"
	MsgMissingAuthHeader      = "missing authorization header"
	MsgInvalidAuthHeader      = "invalid authorization header format"
	MsgInvalidToken           = "invalid or expired token"
	MsgInvalidTokenClaims     = "invalid token claims"
	MsgMissingUserID          = "missing user_id in token"
)
