package repo

import (
	"context"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
)

type FriendRepository interface {
}

type friendRepository struct {
	db *gorm.DB
}

func NewFriendRepository() FriendRepository {
	return &friendRepository{db.DB}
}

func (f *friendRepository) ApplyFriend(ctx context.Context, friend *model.Friend) error {
	friend.Status = "pending"
	friend.CreatedAt = friend.CreatedAt / 1000
	return f.db.WithContext(ctx).Create(friend).Error
}

func (f *friendRepository) PassFriend(ctx context.Context, friendId string) error {
	return f.db.WithContext(ctx).Model(&model.Friend{}).Where("id = ?", friendId).Update("status", "accepted").Error
}

func (f *friendRepository) BlockedFriend(ctx context.Context, friendId string) error {
	return f.db.WithContext(ctx).Model(&model.Friend{}).Where("id = ?", friendId).Update("status", "blocked").Error
}

func (f *friendRepository) DelFriend(ctx context.Context, friendId string) error {
	return f.db.WithContext(ctx).Where("id = ?", friendId).Delete(&model.Friend{}).Error
}
