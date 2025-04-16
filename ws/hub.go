package ws

type Hub struct {
	Clients    map[string]*Client // 所有连接的客户端
	Register   chan *Client       // 客户端上线
	Unregister chan *Client       // 客户端下线
	Broadcast  chan []byte        // 广播消息
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// 注册
		case client := <-h.Register:
			h.Clients[client.UserID] = client
		// 注销
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
			}
		// 广播
		case message := <-h.Broadcast:
			for _, client := range h.Clients {
				client.Send <- message
			}
		}
	}
}
