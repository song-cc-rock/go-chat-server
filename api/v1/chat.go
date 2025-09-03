package v1

import "go-chat-server/internal/model"

type ChatMessage struct {
	ID        string      `json:"id"`
	Send      string      `json:"send"`
	Receiver  string      `json:"receiver"`
	Content   string      `json:"content"`
	CreatedAt int64       `json:"created_at"`
	Avatar    string      `json:"avatar"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	FileId    string      `gorm:"column:file_id" json:"file_id"`
	FileInfo  *model.File `gorm:"foreignKey:FileId;references:ID" json:"fileInfo,omitempty"`
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
