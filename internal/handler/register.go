package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"net/http"
)

type RegisterHandler struct {
	emailService service.EmailService
	userService  service.UserService
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{
		emailService: service.NewEmailService(),
		userService:  service.NewUserService(),
	}
}

func (r *RegisterHandler) SendVerifyCode(ctx *gin.Context) {
	var req v1.SendVerifyCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "Invalid parameter, mail cannot be null")
		return
	}

	if err := r.emailService.SendVerifyCode(req.Mail); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, "Verify code sent successfully")
}

func (r *RegisterHandler) LoginByVerifyCode(ctx *gin.Context) {
	var req v1.LoginByCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "Invalid parameter, mail or code cannot be null")
		return
	}

	if !r.emailService.VerifyCode(req.Mail, req.Code) {
		v1.HandleError(ctx, http.StatusBadRequest, "Invalid verify code")
		return
	}

	token, err := r.userService.GenerateToken(ctx, req.Mail)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	v1.HandleSuccess(ctx, v1.LoginResponse{
		AccessToken: token,
	})
}
