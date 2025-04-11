package v1

type SendVerifyCodeRequest struct {
	Mail string `json:"mail" binding:"required"`
}

type LoginByCodeRequest struct {
	Mail string `json:"mail" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
