package middleware

import (
	"fmt"
	"net/http"

	"test-go/internal/shared/errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RecoveryMiddleware recovers from panics and returns a proper HTTP response
func RecoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace
				logger.WithFields(logrus.Fields{
					"panic":  fmt.Sprintf("%v", err),
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("PANIC RECOVERED")

				// Return error response
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    errors.InternalError,
					"message": errors.ErrInternalError,
					"details": "An unexpected error occurred",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
