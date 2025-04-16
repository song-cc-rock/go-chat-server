package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-chat-server/internal/model"
	"go-chat-server/internal/repo"
	"go-chat-server/pkg/config"
	"go-chat-server/pkg/jwt"
	"net/http"
	"net/url"
)

type GithubService interface {
	GetAuthCodeUrl() (string, error)
	AuthAndGetToken(code string) string
}

type githubService struct {
	userRepo repo.UserRepository
}

func NewGithubService() GithubService {
	return &githubService{
		userRepo: repo.NewUserRepository(),
	}
}

func (a *githubService) GetAuthCodeUrl() (string, error) {
	clientId := config.GetString("oauth.github.client_id")
	redirectUri := config.GetString("oauth.github.redirect_uri")
	if clientId == "" || redirectUri == "" {
		return "", fmt.Errorf("oauth config not found")
	}
	return "client_id=" + clientId + "&redirect_uri=" + redirectUri + "&scope=user", nil
}

func (a *githubService) AuthAndGetToken(code string) string {
	accessToken, err := getAccessToken(code)
	if err != nil {
		return ""
	}
	authUser, err := getAuthUser(accessToken)
	if err != nil {
		return ""
	}

	user := a.getUserInfo(authUser)
	if user == nil {
		return ""
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return ""
	}
	return token
}

func getAccessToken(code string) (*model.Token, error) {
	accessTokenUrl := "https://github.com/login/oauth/access_token"
	data := url.Values{}
	data.Set("client_id", config.GetString("oauth.github.client_id"))
	data.Set("client_secret", config.GetString("oauth.github.client_secret"))
	data.Set("code", code)
	data.Set("redirect_uri", config.GetString("oauth.github.redirect_uri"))

	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodPost, accessTokenUrl, bytes.NewBufferString(data.Encode())); err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	var token model.Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("the code passed is incorrect or expired")
	}
	return &token, nil
}

func getAuthUser(token *model.Token) (map[string]interface{}, error) {
	var userInfoUrl = "https://api.github.com/user"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (a *githubService) getUserInfo(authUser map[string]interface{}) *model.User {
	user, _ := a.userRepo.GetByGithubId(int64(authUser["id"].(float64)))
	if user != nil {
		return user
	}
	githubUser, err := a.userRepo.CreateGithubUser(authUser)
	if err != nil {
		return nil
	}
	return githubUser
}
