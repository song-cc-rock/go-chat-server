package service

import "go-chat-server/internal/repo"

type ChatService interface {
	GetUnReadCount(userId string) (int64, error)
	UpdateMsgFileId(tmpId string, fileId string) error
}

type chatService struct {
	messageRepo repo.MessageRepository
}

func NewChatService() ChatService {
	return &chatService{
		messageRepo: repo.NewMessageRepository(),
	}
}

func (c *chatService) GetUnReadCount(userId string) (int64, error) {
	return c.messageRepo.GetUnReadCount(userId)
}

func (c *chatService) UpdateMsgFileId(tmpId string, fileId string) error {
	return c.messageRepo.UpdateMsgFileId(tmpId, fileId)
}
