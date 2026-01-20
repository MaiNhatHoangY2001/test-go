package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RequestLoggingMiddleware logs all HTTP requests with tracing
func RequestLoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get or generate request ID
		requestID := c.GetString("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
			c.Set("X-Request-ID", requestID)
		}
		c.Header("X-Request-ID", requestID)

		startTime := time.Now()

		// Log request
		logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"query":      c.Request.URL.RawQuery,
			"ip":         c.ClientIP(),
		}).Info("request_started")

		// Process request
		c.Next()

		// Log response
		duration := time.Since(startTime).Milliseconds()
		logger.WithFields(logrus.Fields{
			"request_id":  requestID,
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status":      c.Writer.Status(),
			"duration_ms": duration,
			"bytes":       c.Writer.Size(),
		}).Info("request_completed")
	}
}

// RequestIDMiddleware adds request ID to context
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		c.Set("X-Request-ID", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
