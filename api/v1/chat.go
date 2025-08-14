package v1

type ChatMessage struct {
	ID        string `json:"id"`
	Send      string `json:"send"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	Avatar    string `json:"avatar"`
	Status    string `json:"status"` // 状态: sent, success, failed, deleted
}

type SendMsg struct {
	ID             string `json:"id"`
	ConversationId string `json:"conversation_id"`
	Send           string `json:"send"`
	Receiver       string `json:"receiver"`
	Content        string `json:"content"`
	CreatedAt      int64  `json:"created_at"`
}
