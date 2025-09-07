package service

import (
	"context"
	"go-chat-server/internal/model"
	"go-chat-server/internal/repo"
)

type FriendService interface {
	ApplyFriend(ctx context.Context, friendReq *model.FriendRequest) error
}

type friendService struct {
	friendRepo repo.FriendRepository
}

func NewFriendService() FriendService {
	return &friendService{
		friendRepo: repo.NewFriendRepository(),
	}
}

func (f *friendService) ApplyFriend(ctx context.Context, friendReq *model.FriendRequest) error {
	return f.friendRepo.ApplyFriend(ctx, friendReq)
}
