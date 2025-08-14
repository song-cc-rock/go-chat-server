package service

import (
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/repo"
)

type ConversationService interface {
	GetConversationList(userId string) ([]v1.ConversationResponse, error)
	GetConversationMsgHis(conversationId string) ([]v1.ChatMessage, error)
	ClearConversationUnreadCount(conversationId string) error
}

type conversationService struct {
	conversationRepo repo.ConversationRepository
}

func NewConversationService() ConversationService {
	return &conversationService{
		conversationRepo: repo.NewConversationRepository(),
	}
}

func (c *conversationService) GetConversationList(userId string) ([]v1.ConversationResponse, error) {
	return c.conversationRepo.GetConversationList(userId)
}

func (c *conversationService) GetConversationMsgHis(conversationId string) ([]v1.ChatMessage, error) {
	return c.conversationRepo.GetConversationMsgHis(conversationId)
}

func (c *conversationService) ClearConversationUnreadCount(conversationId string) error {
	return c.conversationRepo.ClearConversationUnreadCount(conversationId)
}
