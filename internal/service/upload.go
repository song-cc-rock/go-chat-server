package service

import (
	"go-chat-server/internal/model"
	"go-chat-server/internal/repo"
)

type FileService interface {
	SaveFileInfo(file *model.File) (string, error)
	GetFileInfo(fileId string) (*model.File, error)
}

type fileService struct {
	fileRepo repo.FileRepository
}

func NewFileService() FileService {
	return &fileService{
		fileRepo: repo.NewFileRepository(),
	}
}

func (u *fileService) SaveFileInfo(file *model.File) (string, error) {
	return u.fileRepo.SaveFileToDB(file)
}

func (u *fileService) GetFileInfo(fileId string) (*model.File, error) {
	return u.fileRepo.GetFileInfo(fileId)
}
