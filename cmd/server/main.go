package main

import (
	"go-chat-server/internal/handler"
	"go-chat-server/pkg/config"
	"go-chat-server/pkg/db"
	"go-chat-server/router"
	"log"
)

func main() {
	// init db
	db.InitDB()

	// start http server
	r := router.Init(handler.NewRegisterHandler())
	addr := config.GetString("server.host") + ":" + config.GetString("server.port")
	_ = r.Run(addr)
	log.Println("Server started at " + addr)
}
