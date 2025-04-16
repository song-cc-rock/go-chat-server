package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-chat-server/pkg/jwt"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServerWs(hub *Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "无法连接Chat服务器, 未认证"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := jwt.ParseToken(tokenStr)
		if err != nil || userId == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "无法连接Chat服务器, 认证失败"})
			return
		}

		// upgrade connection
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "无法连接Chat服务器, 资源异常"})
			return
		}

		// create a new client
		client := &Client{
			UserID: userId,
			Conn:   conn,
			Send:   make(chan []byte, 256),
			Hub:    hub,
		}

		// add client to hub
		hub.Register <- client

		// start read and write pumps
		go client.ReadPump()
		go client.WritePump()
	}
}
