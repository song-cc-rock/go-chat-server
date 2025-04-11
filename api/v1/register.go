package v1

type SendVerifyCodeRequest struct {
	Mail string `json:"mail" binding:"required"`
}

type RegisterByCodeRequest struct {
	Mail     string `json:"mail" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginByPwdRequest struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}
