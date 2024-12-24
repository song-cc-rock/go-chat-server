package api

type Message struct {
	Sender    string `json:"sender"`    // send username
	ClientId  string `json:"clientId"`  // client uuid
	Recipient string `json:"recipient"` // recipient username
	Content   string `json:"content"`
}
