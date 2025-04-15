package handler

import (
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"net/http"
)

type GithubHandler struct {
	githubService service.GithubService
}

func NewGithubHandler() *GithubHandler {
	return &GithubHandler{
		githubService: service.NewGithubService(),
	}
}

func (r *GithubHandler) GetAuthCodeUrl(ctx *gin.Context) {
	url, err := r.githubService.GetAuthCodeUrl()
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	v1.HandleSuccess(ctx, url)
}

func (r *GithubHandler) AuthAndGetToken(ctx *gin.Context) {
	code, _ := ctx.GetQuery("code")
	if code == "" {
		v1.HandleError(ctx, http.StatusBadRequest, "第三方认证失败, 参数异常")
		return
	}
	token := r.githubService.AuthAndGetToken(code)
	if token == "" {
		v1.HandleError(ctx, http.StatusInternalServerError, "获取token失败")
		return
	}

	redirectURL := "http://localhost:8888/auth-success?token=" + token
	ctx.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
