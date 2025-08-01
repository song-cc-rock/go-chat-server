package v1

type ChatMessage struct {
	ID        string `json:"id"`
	Send      string `json:"send"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	Avatar    string `json:"avatar"`
}

type SendMsg struct {
	ConversationId string `json:"conversation_id"`
	Send           string `json:"send"`
	Receiver       string `json:"receiver"`
	Content        string `json:"content"`
	CreatedAt      int64  `json:"created_at"`
}
