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
		Select("conversation.id, conversation.user_id, conversation.target_user_id, conversation.last_message, conversation.unread_count, conversation.last_message_at, user.nick_name, user.avatar").
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
	err := c.db.Model(&model.Message{}).
		Select("id, from_id, to_id, content, created_at").
		Where(
			"(from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?)",
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
