package router

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/internal/handler"
	"go-chat-server/internal/middleware"
	"log"
)

func Init(registerHandler *handler.RegisterHandler) *gin.Engine {
	log.Println("Initializing router...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// User registration and login routes
	r.POST("/send-code", registerHandler.SendVerifyCode)
	r.POST("/register", registerHandler.RegisterNewUser)
	r.POST("/login", registerHandler.LoginByPwd)

	v1 := r.Group("/v1", middleware.JWTAuthMiddleware())
	{
		// User registration and login routes
		v1.POST("/test-token", registerHandler.TestToken)
	}

	return r
}
