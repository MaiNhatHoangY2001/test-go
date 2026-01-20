package middleware

import (
	"fmt"
	"strings"
	"test-go/internal/shared/response"
	"test-go/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims represents JWT token claims
type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// AuthMiddleware validates JWT tokens and extracts user information
func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.UnauthorizedWithCode(ctx, int64(constants.CodeMissingAuthHeader), constants.MsgMissingAuthHeader)
			ctx.Abort()
			return
		}

		// Extract Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.UnauthorizedWithCode(ctx, int64(constants.CodeInvalidAuthHeader), constants.MsgInvalidAuthHeader)
			ctx.Abort()
			return
		}

		// Parse token with claims validation
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
			// Verify algorithm is HMAC-SHA256 (prevent algorithm switching attack)
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})

		// Check if token is valid and extract claims
		if err != nil || !token.Valid {
			response.UnauthorizedWithCode(ctx, int64(constants.CodeInvalidToken), constants.MsgInvalidToken)
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			response.UnauthorizedWithCode(ctx, int64(constants.CodeInvalidTokenClaims), constants.MsgInvalidTokenClaims)
			ctx.Abort()
			return
		}

		// Validate required claims
		if claims.UserID == "" {
			response.UnauthorizedWithCode(ctx, int64(constants.CodeMissingUserID), constants.MsgMissingUserID)
			ctx.Abort()
			return
		}

		// Store user information in context for downstream handlers
		ctx.Set("user_id", claims.UserID)
		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}
