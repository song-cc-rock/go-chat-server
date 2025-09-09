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

// GetApplies 获取当前用户发起的申请列表
func (r *FriendHandler) GetApplies(ctx *gin.Context) {
	var applies []*v1.FriendReqResponse
	applies, err := r.friendService.GetApplies(ctx, ctx.GetString("userId"))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	v1.HandleSuccess(ctx, applies)
}

// GetAccepts 获取当前用户收到的好友申请
func (r *FriendHandler) GetAccepts(ctx *gin.Context) {
	var accepts []*v1.FriendReqResponse
	accepts, err := r.friendService.GetAccepts(ctx, ctx.GetString("userId"))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	v1.HandleSuccess(ctx, accepts)
}
