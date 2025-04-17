package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"net/http"
)

type ChatHandler struct {
	chatService service.ChatService
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		chatService: service.NewChatService(),
	}
}

func (c *ChatHandler) GetUnReadCount(ctx *gin.Context) {
	userId, _ := ctx.GetQuery("id")
	if userId == "" {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}
	count, err := c.chatService.GetUnReadCount(userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "获取未读消息错误")
		return
	}
	v1.HandleSuccess(ctx, count)
}
