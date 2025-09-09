package router

import (
	"github.com/gin-gonic/gin"
	"go-chat-server/internal/handler"
	"go-chat-server/internal/middleware"
	"go-chat-server/ws"
	"log"
)

func Init(registerHandler *handler.RegisterHandler, authHandler *handler.AuthHandler, userHandler *handler.UserHandler,
	chatHandler *handler.ChatHandler, conversationHandler *handler.ConversationHandler, uploadHandler *handler.FileHandler, friendHandler *handler.FriendHandler,
	hub *ws.Hub) *gin.Engine {
	log.Println("🚀 Initializing router...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// User registration and login routes
	r.POST("/send-code", registerHandler.SendVerifyCode)
	r.POST("/register", registerHandler.RegisterNewUser)
	r.POST("/login", registerHandler.LoginByPwd)
	r.GET("/github/auth-url", authHandler.GetGithubAuthCodeUrl)
	r.GET("/oauth/github", authHandler.AuthGithubAndGetToken)

	v1 := r.Group("/v1", middleware.JWTAuthMiddleware())
	{
		// User registration and login routes
		v1.GET("/user-profile", authHandler.GetAuthUserProfile)
		v1.GET("/user-search", userHandler.SearchByKeyword)
		v1.GET("/unread-count", chatHandler.GetUnReadCount)
		v1.GET("/conversations", conversationHandler.GetConversationList)
		v1.GET("/conversation/his", conversationHandler.GetConversationMsgHis)
		v1.GET("/conversation/clear", conversationHandler.ClearConversationUnreadCount)
		v1.POST("/upload", uploadHandler.Upload)
		v1.GET("/download", uploadHandler.Download)
		v1.POST("/friend/apply", friendHandler.ApplyFriend)
		v1.GET("/applies", friendHandler.GetApplies)
	}

	// Chat socket routes
	r.GET("/ws", ws.ServerWs(hub))

	// Static
	r.Static("/static", "./static")
	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	return r
}
