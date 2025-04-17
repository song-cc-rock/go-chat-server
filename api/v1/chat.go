package v1

type ChatMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Time    int64  `json:"time"`
}
