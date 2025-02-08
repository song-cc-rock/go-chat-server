package main

import (
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"go-chat-server/db"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}

	// client map pool
	globalClients = make(map[int16]*websocket.Conn)

	// msg channel
	//msgChan = make(chan api.Message)

	// client unique key
	clientKey int16 = 1
)
var appViper = viper.New()

// keep client connection
func startConn(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to connect server: %+v", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error close conn: %v", err)
		}
	}(conn)
	addGlobalConn(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error read message: %v", err)
			break
		}
		conn, content := parseMsg(message)
		if conn != nil {
			sendMessage(conn, content)
		} else {
			log.Printf("Msg format error: %s", message)
		}
	}
}

func addGlobalConn(conn *websocket.Conn) {
	globalClients[clientKey] = conn
	log.Printf("Client %d Join", clientKey)
	clientKey++
}

func parseMsg(msg []byte) (*websocket.Conn, string) {
	// msg format: client_key|content
	msgArr := strings.SplitN(string(msg), ":", 2)
	if len(msgArr) != 2 {
		return nil, ""
	} else {
		clientKey, err := strconv.ParseInt(msgArr[0], 10, 16)
		if err != nil {
			log.Printf("Uknown client: %v", err)
			return nil, ""
		}
		if globalClients[int16(clientKey)] != nil {
			return globalClients[int16(clientKey)], msgArr[1]
		} else {
			return nil, ""
		}
	}
}

// send message to target client
func sendMessage(conn *websocket.Conn, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Send msg to client error:", err)
	}
	log.Printf("Send {%s} to client success", message)
}

func main() {
	// init viper config
	viperPath, err := os.Getwd()
	if err != nil {
		panic("Cannot get project path")
	}
	appViper.AddConfigPath(viperPath + string(os.PathSeparator) + "config")
	appViper.SetConfigName("app")
	appViper.SetConfigType("yaml")

	if err := appViper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	} else {
		log.Printf("Using config file: %s", appViper.ConfigFileUsed())
	}

	// init db
	db.InitDB(appViper)

	// start server
	http.HandleFunc("/ws", startConn)
	log.Printf("Websocket server started on 127.0.0.1:8080")
	_ = http.ListenAndServe(":8080", nil)
}
