package repo

import (
	"github.com/google/uuid"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
)

type FileRepository interface {
	SaveFileToDB(file *model.File) (string, error)
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository() FileRepository {
	return &fileRepository{db.DB}
}

func (f *fileRepository) SaveFileToDB(file *model.File) (string, error) {
	// 插入消息
	file.ID = uuid.NewString()
	if err := f.db.Create(file).Error; err != nil {
		return "", err
	}
	// 文件ID
	return file.ID, nil
}
