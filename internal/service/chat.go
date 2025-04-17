package service

import "go-chat-server/internal/repo"

type ChatService interface {
	GetUnReadCount(userId string) (map[string]int64, error)
}

type chatService struct {
	messageRepo repo.MessageRepository
}

func NewChatService() ChatService {
	return &chatService{
		messageRepo: repo.NewMessageRepository(),
	}
}

func (c *chatService) GetUnReadCount(userId string) (map[string]int64, error) {
	return c.messageRepo.GetUnReadCount(userId)
}
