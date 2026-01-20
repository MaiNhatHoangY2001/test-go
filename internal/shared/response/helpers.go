package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnauthorizedWithCode sends a 401 Unauthorized response with specific error code
func UnauthorizedWithCode(c *gin.Context, errorCode int64, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
	})
}

// BadRequestWithCode sends a 400 Bad Request response with specific error code
func BadRequestWithCode(c *gin.Context, errorCode int64, message string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
	})
}

// NotFoundWithCode sends a 404 Not Found response with specific error code
func NotFoundWithCode(c *gin.Context, errorCode int64, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
	})
}

// ConflictWithCode sends a 409 Conflict response with specific error code
func ConflictWithCode(c *gin.Context, errorCode int64, message string) {
	c.JSON(http.StatusConflict, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
	})
}

// InternalServerErrorWithCode sends a 500 Internal Server Error response with specific error code
func InternalServerErrorWithCode(c *gin.Context, errorCode int64, message string) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
	})
}
