package repo

import (
	"github.com/google/uuid"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMsgToDB(message *v1.ChatMessage) (*model.Message, error)
	UpdateMsgStatus(msgIds []string, newStatus string) error
	GetUnReadCount(userId string) (int64, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{db.DB}
}

// SaveMsgToDB 消息入库
func (m *messageRepository) SaveMsgToDB(message *v1.ChatMessage) (*model.Message, error) {
	msg := &model.Message{
		ID:        uuid.NewString(),
		FromId:    message.From,
		ToId:      message.To,
		Content:   message.Content,
		MsgType:   "text",
		Status:    "sent",
		CreatedAt: message.CreatedAt,
	}
	if err := m.db.Create(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

// UpdateMsgStatus 更新消息状态
func (m *messageRepository) UpdateMsgStatus(msgIds []string, newStatus string) error {
	if err := m.db.Model(&model.Message{}).Where("id in ?", msgIds).Update("status", newStatus).Error; err != nil {
		return err
	}
	return nil
}

// GetUnReadCount 获取未读消息数量
func (m *messageRepository) GetUnReadCount(userId string) (int64, error) {
	var count int64
	if err := m.db.Model(&model.Message{}).Select("count(*) as count").Where("to_id = ? and status = ?", userId, "sent").Find(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
