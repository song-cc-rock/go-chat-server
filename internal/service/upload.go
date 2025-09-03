package service

import "go-chat-server/internal/repo"

type UploadService interface {
}

type uploadService struct {
	messageRepo repo.MessageRepository
}

func NewUploadService() UploadService {
	return &uploadService{
		messageRepo: repo.NewMessageRepository(),
	}
}
