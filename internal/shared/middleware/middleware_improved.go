package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RequestIDMiddleware generates a unique request ID for tracking
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Set("X-Request-ID", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// RequestLoggingMiddleware logs request and response details
func RequestLoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		
		requestID := c.GetString("X-Request-ID")

		// Log request
		logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Debug("request_started")

		// Continue to next handler
		c.Next()

		// Log response
		duration := time.Since(startTime)
		logger.WithFields(logrus.Fields{
			"request_id":   requestID,
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"status_code":  c.Writer.Status(),
			"response_time_ms": duration.Milliseconds(),
			"bytes_written": c.Writer.Size(),
		}).Info("request_completed")
	}
}

// RecoveryMiddleware handles panics gracefully
func RecoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID := c.GetString("X-Request-ID")
				logger.WithFields(logrus.Fields{
					"request_id": requestID,
					"error":      err,
					"path":       c.Request.URL.Path,
					"method":     c.Request.Method,
				}).Error("panic_recovered")

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
					"code":  "INTERNAL_ERROR",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// TimeoutMiddleware adds request timeout
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := gin.DefaultContext()
		defer cancel()

		done := make(chan struct{})
		
		go func() {
			c.Next()
			done <- struct{}{}
		}()

		select {
		case <-done:
			// Request completed
		case <-time.After(timeout):
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "request timeout",
				"code":  "REQUEST_TIMEOUT",
			})
			c.Abort()
		}
	}
}

// RateLimitMiddleware placeholder for rate limiting
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement rate limiting (e.g., using redis or token bucket)
		c.Next()
	}
}

// generateRequestID creates a unique request identifier
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randString(8)
}

// randString generates random string for request IDs
func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[0]
	}
	return string(b)
}
