package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"go-chat-server/internal/model"
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

		var msg model.Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			continue
		}

		// TODO: 消息入库

		//c.Hub.Broadcast <- msgBytes
		if targetClient, ok := c.Hub.Clients[msg.To]; ok {
			targetClient.Send <- msgBytes
		} else {
			// TODO: 目标用户不在线，消息入库
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
