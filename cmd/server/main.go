package main

import (
	"go-chat-server/internal/handler"
	"go-chat-server/pkg/bootstrap"
	"go-chat-server/pkg/config"
	"go-chat-server/router"
	"go-chat-server/ws"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	// init db
	if err := bootstrap.Init(); err != nil {
		log.Fatal("❌ Init error:", err.Error())
	}

	// init chat hub
	hub := ws.NewHub()
	go hub.Run()

	// start http server or ws server
	r := router.Init(handler.NewRegisterHandler(), handler.NewAuthHandler(), handler.NewUserHandler(), handler.NewChatHandler(),
		handler.NewConversationHandler(), handler.NewUploadHandler(), handler.NewFriendHandler(), hub)
	addr := config.GetString("server.host") + ":" + config.GetString("server.port")
	log.Println("🚀 Server started at " + addr)
	_ = r.Run(addr)
}
