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
	UpdateConversationLastInfo(message *v1.SendMsg)
	ClearConversationUnreadCount(conversationId string) error
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

	// 更新对话未读消息数 = 0
	c.db.Table("conversation").
		Where("id = ?", conversationId).
		Update("unread_count", 0)

	// 获取会话的消息
	var messages []v1.ChatMessage
	err := c.db.Table("message").
		Select("message.id, u1.id as send, u2.id as receiver, content, created_at, u1.avatar as avatar, message.status").
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

func (c *conversationRepository) UpdateConversationLastInfo(msg *v1.SendMsg) {
	// 更新发送人和接收人的会话信息
	c.db.Table("conversation").
		Where("user_id = ? AND target_user_id = ?", msg.Send, msg.Receiver).
		Updates(map[string]interface{}{
			"last_message_at": msg.CreatedAt / 1000,
			"last_message":    msg.Content,
			"last_sent_user":  msg.Send,
			"unread_count":    0,
		})

	c.db.Table("conversation").
		Where("user_id = ? AND target_user_id = ?", msg.Receiver, msg.Send).
		Updates(map[string]interface{}{
			"last_message_at": msg.CreatedAt / 1000,
			"last_message":    msg.Content,
			"last_sent_user":  msg.Send,
			"unread_count":    gorm.Expr("unread_count + 1"),
		})

}

func (c *conversationRepository) ClearConversationUnreadCount(conversationId string) error {
	return c.db.Table("conversation").
		Where("id = ?", conversationId).
		Update("unread_count", 0).Error
}
