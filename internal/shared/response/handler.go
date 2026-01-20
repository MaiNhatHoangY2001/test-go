package response

import (
	"net/http"

	errs "test-go/internal/shared/errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// HandleError handles errors with proper HTTP status codes and logging
func HandleError(c *gin.Context, logger *logrus.Logger, err error) {
	if appErr, ok := err.(*errs.AppError); ok {
		statusCode := appErr.HTTPStatusCode()

		logger.WithFields(logrus.Fields{
			"error_code": appErr.Code,
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
			"status":     statusCode,
		}).Warn(appErr.Error())

		c.JSON(statusCode, ErrorResponse{
			Code:    string(appErr.Code),
			Message: appErr.Message,
			Details: appErr.Details,
		})
		return
	}

	// Generic error - treat as internal
	logger.WithFields(logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
		"status": http.StatusInternalServerError,
	}).Error(err.Error())

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    string(errs.InternalError),
		Message: errs.ErrInternalError,
		Details: err.Error(),
	})
}

// Success sends a successful JSON response
func Success(c *gin.Context, data interface{}, statusCode int) {
	c.JSON(statusCode, SuccessResponse{Data: data})
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

// BadRequest sends a 400 Bad Request response
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Code:    string(errs.BadRequestError),
		Message: message,
	})
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Code:    string(errs.NotFoundError),
		Message: message,
	})
}

// Conflict sends a 409 Conflict response
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, ErrorResponse{
		Code:    string(errs.ConflictError),
		Message: message,
	})
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Code:    string(errs.UnauthorizedError),
		Message: message,
	})
}
