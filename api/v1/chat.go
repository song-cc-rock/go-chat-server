package v1

type ChatMessage struct {
	ID        string `json:"id"`
	Send      string `json:"send"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	Avatar    string `json:"avatar"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	//FileInfo  *model.File `json:"fileInfo,omitempty"`
}

type SendMsg struct {
	ID             string `json:"id"`
	ConversationId string `json:"conversation_id"`
	Send           string `json:"send"`
	Receiver       string `json:"receiver"`
	Content        string `json:"content"`
	CreatedAt      int64  `json:"created_at"`
	Type           string `json:"type"`
}
