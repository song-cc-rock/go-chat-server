package service

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"go-chat-server/pkg/config"
	"go-chat-server/pkg/utils"
	"time"
)

var (
	stateCode = cache.New(1*time.Minute, 5*time.Minute)
)

type AuthService interface {
	GetAuthCodeUrl(authType string) (string, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (a *authService) GetAuthCodeUrl(authType string) (string, error) {
	clientId := config.GetString("oauth." + authType + ".client_id")
	redirectUri := config.GetString("oauth." + authType + ".redirect_uri")
	state := utils.GenerateVerifyCode()
	if clientId == "" || redirectUri == "" || state == "" {
		return "", fmt.Errorf("oauth config not found")
	}
	stateCode.Set(authType, state, cache.DefaultExpiration)
	return "client_id=" + clientId + "&redirect_uri=" + redirectUri + "&state=" + state, nil
}
