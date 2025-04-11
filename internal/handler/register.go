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

	v1.HandleSuccess(ctx, "Verify code send successfully")
}

func (r *RegisterHandler) RegisterNewUser(ctx *gin.Context) {
	var req v1.RegisterByCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "Invalid parameter")
		return
	}

	if !r.userService.IsNewUser(ctx, req.Mail) {
		v1.HandleError(ctx, http.StatusBadRequest, "This email is already registered")
		return
	}

	if !r.emailService.VerifyCode(req.Mail, req.Code) {
		v1.HandleError(ctx, http.StatusBadRequest, "Invalid verify code")
		return
	}

	user, err := r.userService.RegisterNewUser(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, user.NickName)
}

func (r *RegisterHandler) LoginByPwd(ctx *gin.Context) {
	var req v1.LoginByPwdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "Invalid parameter, mail or password cannot be null")
		return
	}

	token := r.userService.VerifyPwdWithToken(ctx, &req)
	if token == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	v1.HandleSuccess(ctx, v1.LoginResponse{
		AccessToken: token,
	})
}
