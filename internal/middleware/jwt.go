package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"go-chat-server/config"
	"time"
)

type Claims struct {
	UserId string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(userId string) (string, error) {
	expired := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expired),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-chat-server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetString("jwt.secret"))
}
