package router

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/internal/handler"
	"log"
)

func Init(registerHandler *handler.RegisterHandler) *gin.Engine {
	log.Println("Initializing router...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		// User registration and login routes
		v1.POST("/send-code", registerHandler.SendVerifyCode)
		v1.POST("/register", registerHandler.RegisterNewUser)
		v1.POST("/login", registerHandler.LoginByPwd)
	}

	return r
}
