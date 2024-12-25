package test

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"go-chat-server/api"
	"log"
	"os"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %+v", err)
		return
	}

	defer conn.Close()
	go func() {
		for {
			var msg api.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("Error reading message: ", err)
				return
			}
			fmt.Printf("Message from server: %+v\n", msg)
		}
	}()

	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		msg := api.Message{Sender: name, Content: message}
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
}
