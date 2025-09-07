package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/internal/service"
	"net/http"
)

type FriendHandler struct {
	friendService service.FriendService
}

func NewFriendHandler() *FriendHandler {
	return &FriendHandler{
		friendService: service.NewFriendService(),
	}
}

func (r *FriendHandler) ApplyFriend(ctx *gin.Context) {
	var friendReq model.FriendRequest
	if err := ctx.ShouldBindJSON(&friendReq); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}
	if err := r.friendService.ApplyFriend(ctx, &friendReq); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	v1.HandleSuccess(ctx, "success")
}
