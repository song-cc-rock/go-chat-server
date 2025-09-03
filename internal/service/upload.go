package service

import (
	"go-chat-server/internal/model"
	"go-chat-server/internal/repo"
)

type UploadService interface {
	SaveFileInfo(file *model.File) (string, error)
}

type uploadService struct {
	fileRepo repo.FileRepository
}

func NewUploadService() UploadService {
	return &uploadService{
		fileRepo: repo.NewFileRepository(),
	}
}

func (u *uploadService) SaveFileInfo(file *model.File) (string, error) {
	return u.fileRepo.SaveFileToDB(file)
}
