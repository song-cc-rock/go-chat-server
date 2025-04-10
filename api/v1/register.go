package v1

type SendVerifyCodeRequest struct {
	Mail string `json:"mail" binding:"required,mail"`
}

type LoginByCodeRequest struct {
	Mail string `json:"mail" binding:"required,mail"`
	Code string `json:"code" binding:"required,code"`
}

type LoginResponse struct {
	AccessToken string `json:"token"`
}
