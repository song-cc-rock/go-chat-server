package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"net/http"
)

type ConversationHandler struct {
	conversationService service.ConversationService
}

func NewConversationHandler() *ConversationHandler {
	return &ConversationHandler{
		conversationService: service.NewConversationService(),
	}
}

func (c *ConversationHandler) GetConversationList(ctx *gin.Context) {
	userId, _ := ctx.GetQuery("id")
	if userId == "" {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}
	list, err := c.conversationService.GetConversationList(userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "获取会话列表错误")
		return
	}
	v1.HandleSuccess(ctx, list)
}
