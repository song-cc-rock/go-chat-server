package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (r *AuthHandler) GetAuthCodeUrl(ctx *gin.Context) {
	var req v1.AuthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}

	url, err := r.authService.GetAuthCodeUrl(req.AuthType)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, url)
}
