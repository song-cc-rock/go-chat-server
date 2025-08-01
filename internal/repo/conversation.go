package repo

import (
	"fmt"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	GetConversationList(userId string) ([]v1.ConversationResponse, error)
	GetConversationMsgHis(conversationId string) ([]v1.ChatMessage, error)
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository() ConversationRepository {
	return &conversationRepository{
		db: db.DB,
	}
}

// GetConversationList 获取会话列表
func (c *conversationRepository) GetConversationList(userId string) ([]v1.ConversationResponse, error) {
	var conversations []v1.ConversationResponse
	err := c.db.Table("conversation").
		Select("conversation.id, conversation.user_id, conversation.target_user_id, conversation.last_message, conversation.unread_count, conversation.last_message_at,"+
			"conversation.last_sent_user, user.nick_name, user.avatar").
		Joins("JOIN user ON user.id = conversation.target_user_id").
		Where("conversation.user_id = ?", userId).
		Find(&conversations).Error
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

// GetConversationMsgHis 获取会话历史消息
func (c *conversationRepository) GetConversationMsgHis(conversationId string) ([]v1.ChatMessage, error) {
	conversation := &model.Conversation{}
	if err := c.db.Where("id = ?", conversationId).First(conversation).Error; err != nil {
		return nil, fmt.Errorf("failed to get conversation msg: %v", err)
	}
	var messages []v1.ChatMessage
	err := c.db.Table("message").
		Select("message.id, u1.nick_name as send, u2.nick_name as send, content, created_at, u1.avatar as avatar").
		Joins("join user u1 on u1.id = message.from_id").
		Joins("join user u2 on u2.id = message.to_id").
		Where(
			"(message.from_id = ? and message.to_id = ?) or (message.from_id = ? and message.to_id = ?)",
			conversation.UserID, conversation.TargetUserID,
			conversation.TargetUserID, conversation.UserID,
		).
		Order("created_at asc").
		Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
