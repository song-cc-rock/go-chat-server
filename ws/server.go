package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServerWs(hub *Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// upgrade connection
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "无法连接Chat服务器, 资源异常"})
			return
		}

		// create a new client
		client := &Client{
			Conn: conn,
			Send: make(chan []byte, 256),
			Hub:  hub,
		}

		// add client to hub
		hub.Register <- client

		// TODO: 新上线的用户, 需要将MQ中消息推送

		// start read and write pumps
		go client.ReadPump()
		go client.WritePump()
	}
}
