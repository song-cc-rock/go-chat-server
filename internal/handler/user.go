package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// SearchByKeyword 根据关键词精准匹配用户
func (u *UserHandler) SearchByKeyword(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	if keyword == "" {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常, 搜索关键字不能为空")
	}
	currentUserId, _ := ctx.Get("userId")
	user, err := u.userService.GetUserByKeyword(ctx, keyword, currentUserId.(string))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "搜索用户失败: "+err.Error())
		return
	}
	v1.HandleSuccess(ctx, user)
}
