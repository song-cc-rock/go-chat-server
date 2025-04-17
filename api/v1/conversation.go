package v1

type ConversationResponse struct {
	ID            string `json:"id"`
	TargetUserID  string `json:"target_user_id"`
	NickName      string `json:"nick_name"`
	Avatar        string `json:"avatar"`
	LastMessage   string `json:"last_message"`
	LastMessageAt int64  `json:"last_message_at"`
	UnreadCount   int64  `json:"unread_count"`
}
