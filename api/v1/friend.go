package v1

type FriendReqResponse struct {
	ID        string `json:"id"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"createdAt"`
}
