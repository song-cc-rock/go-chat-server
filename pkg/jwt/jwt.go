package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"go-chat-server/pkg/config"
	"time"
)

type Claims struct {
	UserId string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (string, error) {
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
	return token.SignedString([]byte(config.GetString("jwt.secret")))
}

func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.GetString("jwt.secret"), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
