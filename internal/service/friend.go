package service

import (
	"context"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/internal/repo"
)

type FriendService interface {
	ApplyFriend(ctx context.Context, friendReq *model.FriendRequest) error
	GetApplies(ctx context.Context, fromId string) ([]*v1.FriendReqResponse, error)
	GetAccepts(ctx context.Context, fromId string) ([]*v1.FriendReqResponse, error)
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

func (f *friendService) GetApplies(ctx context.Context, fromId string) ([]*v1.FriendReqResponse, error) {
	return f.friendRepo.GetApplies(ctx, fromId)
}

func (f *friendService) GetAccepts(ctx context.Context, fromId string) ([]*v1.FriendReqResponse, error) {
	return f.friendRepo.GetAccepts(ctx, fromId)
}
