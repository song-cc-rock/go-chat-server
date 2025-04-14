package v1

type AuthRequest struct {
	AuthType string `json:"authType" binding:"required"`
}
