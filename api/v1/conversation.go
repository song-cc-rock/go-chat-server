package v1

type ConversationResponse struct {
	ID            string `json:"id"`
	TargetUserID  string `json:"targetUserId"`
	NickName      string `json:"nickName"`
	Avatar        string `json:"avatar"`
	LastMessage   string `json:"lastMessage"`
	LastMessageAt int64  `json:"lastMessageAt"`
	LastSentUser  string `json:"lastSentUser"`
	UnreadCount   int64  `json:"unreadCount"`
}
