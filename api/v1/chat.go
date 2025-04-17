package v1

type ChatMessage struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}
