package response

import (
	"net/http"

	errs "test-go/internal/shared/errors"
	"test-go/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Legacy error response - keeping for compatibility
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Legacy success response - keeping for compatibility
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// HandleError handles errors with proper HTTP status codes and logging using new standard format
func HandleError(c *gin.Context, logger *logrus.Logger, err error) {
	if appErr, ok := err.(*errs.AppError); ok {
		statusCode := appErr.HTTPStatusCode()

		logger.WithFields(logrus.Fields{
			"error_code": appErr.Code,
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
			"status":     statusCode,
		}).Warn(appErr.Error())

		// Map error type to unique error code
		var errorCode int64
		switch appErr.Code {
		case errs.UnauthorizedError:
			errorCode = int64(constants.CodeUnauthorized)
		case errs.BadRequestError:
			errorCode = int64(constants.CodeBadRequest)
		case errs.NotFoundError:
			errorCode = int64(constants.CodeNotFound)
		case errs.ConflictError:
			errorCode = int64(constants.CodeConflict)
		case errs.ValidationError:
			errorCode = int64(constants.CodeValidationError)
		default:
			errorCode = int64(constants.CodeInternalError)
		}

		c.JSON(statusCode, APIResponse{
			Success: false,
			Error: &ErrorInfo{
				Code:    errorCode,
				Message: appErr.Message,
			},
		})
		return
	}

	// Generic error - treat as internal
	logger.WithFields(logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
		"status": http.StatusInternalServerError,
	}).Error(err.Error())

	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    int64(constants.CodeInternalError),
			Message: errs.ErrInternalError,
		},
	})
}

// Success sends a successful JSON response using new standard format
func Success(c *gin.Context, data interface{}, statusCode int) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Data:    data,
	})
}

// SuccessPaging sends a successful paging JSON response using new standard format
func SuccessPaging(c *gin.Context, page, pageSize, totalElement int64, data interface{}) {
	c.JSON(http.StatusOK, PagingResponse{
		Success: true,
		Data: &PagingData{
			Page:         page,
			PageSize:     pageSize,
			TotalElement: totalElement,
			Data:         data,
		},
	})
}

// Created sends a 201 Created response
func Created(c *gin.Context, data interface{}) {
	Success(c, data, http.StatusCreated)
}

// OK sends a 200 OK response
func OK(c *gin.Context, data interface{}) {
	Success(c, data, http.StatusOK)
}

// NoContent sends a 204 No Content response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// BadRequest sends a 400 Bad Request response using new standard format
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    int64(constants.CodeBadRequest),
			Message: message,
		},
	})
}

// NotFound sends a 404 Not Found response using new standard format
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    int64(constants.CodeNotFound),
			Message: message,
		},
	})
}

// Conflict sends a 409 Conflict response using new standard format
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    int64(constants.CodeConflict),
			Message: message,
		},
	})
}

// Unauthorized sends a 401 Unauthorized response using new standard format
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    int64(constants.CodeUnauthorized),
			Message: message,
		},
	})
}

// InternalServerError sends a 500 Internal Server Error response using new standard format
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    int64(constants.CodeInternalError),
			Message: message,
		},
	})
}
