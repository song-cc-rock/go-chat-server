package main

import (
	"go-chat-server/internal/handler"
	"go-chat-server/pkg/config"
	"go-chat-server/pkg/db"
	"go-chat-server/router"
	"go-chat-server/ws"
	"log"
)

func main() {
	// init db
	db.InitDB()

	// init chat hub
	hub := ws.NewHub()
	go hub.Run()

	// start http server or ws server
	r := router.Init(handler.NewRegisterHandler(), handler.NewAuthHandler(), handler.NewChatHandler(), hub)
	addr := config.GetString("server.host") + ":" + config.GetString("server.port")
	_ = r.Run(addr)

	log.Println("Server started at " + addr)
}
