package repo

import (
	"github.com/google/uuid"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMsgToDB(message *v1.SendMsg) (string, error)
	UpdateMsgStatus(msgIds []string, newStatus string) error
	GetUnReadCount(userId string) (int64, error)
	UpdateMsgFileId(tmpId string, fileId string) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{db.DB}
}

// SaveMsgToDB 消息入库
func (m *messageRepository) SaveMsgToDB(message *v1.SendMsg) (string, error) {
	// 插入消息
	msg := &model.Message{
		ID:        uuid.NewString(),
		FromId:    message.Send,
		ToId:      message.Receiver,
		Content:   message.Content,
		MsgType:   message.Type,
		Status:    "sent",
		CreatedAt: message.CreatedAt / 1000,
		FileId:    message.ID,
	}
	if err := m.db.Create(msg).Error; err != nil {
		return "", err
	}
	// 更新对话字段
	return msg.ID, nil
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
	if err := m.db.Model(&model.Conversation{}).
		Select("ifnull(sum(unread_count), 0) as count").
		Where("user_id = ?", userId).Find(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// UpdateMsgFileId 更新实际文件ID
func (m *messageRepository) UpdateMsgFileId(tmpId string, fileId string) error {
	if err := m.db.Model(&model.Message{}).Where("file_id = ?", tmpId).Update("file_id", fileId).Error; err != nil {
		return err
	}
	return nil
}
