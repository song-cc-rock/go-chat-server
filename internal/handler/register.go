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
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}

	if err := r.emailService.SendVerifyCode(req.Mail); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, "success")
}

func (r *RegisterHandler) RegisterNewUser(ctx *gin.Context) {
	var req v1.RegisterByCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}

	if !r.userService.IsNewUser(ctx, req.Mail) {
		v1.HandleError(ctx, http.StatusBadRequest, "邮箱已注册, 请直接登录")
		return
	}

	if !r.emailService.VerifyCode(req.Mail, req.Code) {
		v1.HandleError(ctx, http.StatusBadRequest, "验证码错误或过期")
		return
	}

	user, err := r.userService.RegisterNewUser(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, user.NickName)
}

func (r *RegisterHandler) UpdatePwd(ctx *gin.Context) {
	var req v1.RegisterByCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}

	if r.userService.IsNewUser(ctx, req.Mail) {
		v1.HandleError(ctx, http.StatusBadRequest, "邮箱未注册, 请先注册")
		return
	}

	if !r.emailService.VerifyCode(req.Mail, req.Code) {
		v1.HandleError(ctx, http.StatusBadRequest, "验证码错误或过期")
		return
	}

	update, err := r.userService.UpdatePwd(ctx, req.Mail, req.Password)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, update)
}

func (r *RegisterHandler) LoginByPwd(ctx *gin.Context) {
	var req v1.LoginByPwdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, "参数异常")
		return
	}

	token := r.userService.VerifyPwdWithToken(ctx, &req)
	if token == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, "账号或密码错误")
		return
	}

	v1.HandleSuccess(ctx, token)
}
