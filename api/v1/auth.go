package v1

type AuthRequest struct {
	AuthType string `json:"authType" binding:"required"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type AuthUserResponse struct {
	ID       string `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}
