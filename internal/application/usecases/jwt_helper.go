package usecases

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func SetJWTSecretForUsecases(secret []byte) {
	jwtSecret = secret
}

func GenerateJWT(email, userID string) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
