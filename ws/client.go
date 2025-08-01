package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/repo"
	"go-chat-server/pkg/jwt"
)

type Client struct {
	UserID string
	Conn   *websocket.Conn
	Send   chan []byte
	Hub    *Hub
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		err := c.Conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		_, msgBytes, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			continue
		}

		if msg["type"] == "auth" {
			token := msg["token"].(string)
			userId, err := jwt.ParseToken(token)
			if err != nil || userId == "" {
				err := c.Conn.WriteMessage(websocket.TextMessage, []byte("auth failed"))
				if err != nil {
					return
				}
				c.Conn.Close()
				return
			}
			c.UserID = userId
			// auth success
			continue
		}

		// no auth
		if c.UserID == "" {
			err := c.Conn.WriteMessage(websocket.TextMessage, []byte("auth failed, cannot send message"))
			if err != nil {
				return
			}
			continue
		}

		// decode message
		var sendMsg v1.SendMsg
		mapstructure.Decode(msg, &sendMsg)
		msgRepo := repo.NewMessageRepository()
		_, err = msgRepo.SaveMsgToDB(&sendMsg)
		if err != nil {
			continue
		}

		//c.Hub.Broadcast <- msgBytes
		if targetClient, ok := c.Hub.Clients[sendMsg.Receiver]; ok {
			targetClient.Send <- msgBytes
			// TODO: 不应该在这里更新消息状态, 需要用户在客户端对话框确认
		} else {
			// TODO: 目标用户不在线，入MQ
		}
	}
}

func (c *Client) WritePump() {
	defer func(Conn *websocket.Conn) {
		err := Conn.Close()
		if err != nil {

		}
	}(c.Conn)
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				return
			}
			err := c.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}
