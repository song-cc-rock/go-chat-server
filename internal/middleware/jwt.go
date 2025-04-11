package middleware

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/pkg/jwt"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "no token provided"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := jwt.ParseToken(tokenStr)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			return
		}

		ctx.Next()
	}
}
