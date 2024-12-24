package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go-chat-server/api"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client represents a connected client
type Client struct {
	conn *websocket.Conn
	uuid uuid.UUID // unique id
}

// client map pool
var clients = make(map[*Client]bool)

// msg channel
var msgChan = make(chan api.Message)

func startConn(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to connect server: %+v", err)
		return
	}
	clientId, err := uuid.NewRandom()
	defer conn.Close()
	client := &Client{conn: conn, uuid: clientId}
	clients[client] = true
	log.Printf("Client %s connected", client.uuid.String())

	for {
		var msg api.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, client)
			break
		}
		msgChan <- msg
	}
}

// sends message to all clients
func sendMsg() {
	for {
		msg := <-msgChan
		for client := range clients {
			err := client.conn.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing message to %s : %v", client.uuid.String(), err)
				client.conn.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", startConn)
	go sendMsg()
	log.Printf("Websocket server started on localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}
