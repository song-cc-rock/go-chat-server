package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/repo"
	"go-chat-server/pkg/jwt"
	"time"
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
			c.Hub.Register <- c
			// auth success and register to hub
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
		ackMsg := map[string]interface{}{
			"clientTmpId": sendMsg.ID,
			"type":        "ack",
			"timestamp":   time.Now().UnixMilli(),
		}
		msgRepo := repo.NewMessageRepository()
		dbMsgId := ""
		dbMsgId, err = msgRepo.SaveMsgToDB(&sendMsg)
		if err != nil {
			ackMsg["status"] = "failed"
			continue
		}

		ackMsg["status"] = "success"
		ackMsg["actualId"] = dbMsgId
		// update message status in db
		_ = msgRepo.UpdateMsgStatus([]string{dbMsgId}, "success")
		conversationRepo := repo.NewConversationRepository()
		conversationRepo.UpdateConversationLastInfo(&sendMsg)

		//c.Hub.Broadcast <- msgBytes
		if targetClient, ok := c.Hub.Clients[sendMsg.Receiver]; ok {
			var msg map[string]interface{}
			if err := json.Unmarshal(msgBytes, &msg); err != nil {
				continue
			}
			msg["id"] = dbMsgId
			msg["status"] = "success"
			sendBytes, _ := json.Marshal(msg)
			targetClient.Send <- sendBytes
		} else {
			// TODO: target user not online, send to hub
		}

		// send ack message to sender
		ackBytes, _ := json.Marshal(ackMsg)
		c.Send <- ackBytes
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
