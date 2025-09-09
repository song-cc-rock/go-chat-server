package repo

import (
	"context"
	"github.com/google/uuid"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
)

type FriendRepository interface {
	ApplyFriend(ctx context.Context, friendReq *model.FriendRequest) error
	GetApplies(ctx context.Context, fromId string) ([]*v1.FriendReqResponse, error)
}

type friendRepository struct {
	db *gorm.DB
}

func NewFriendRepository() FriendRepository {
	return &friendRepository{db.DB}
}

func (f *friendRepository) ApplyFriend(ctx context.Context, friendReq *model.FriendRequest) error {
	friendReq.ID = uuid.NewString()
	friendReq.Status = "pending"
	friendReq.CreatedAt = friendReq.CreatedAt / 1000
	return f.db.WithContext(ctx).Model(&model.FriendRequest{}).Create(friendReq).Error
}

// GetApplies 获取当前用户的申请列表
func (f *friendRepository) GetApplies(ctx context.Context, fromId string) ([]*v1.FriendReqResponse, error) {
	var friendReqs []*v1.FriendReqResponse
	if err := f.db.WithContext(ctx).Model(&model.FriendRequest{}).
		Select("friend_request.id, user.avatar, user.nick_name as name, friend_request.message, friend_request.status, friend_request.created_at").
		Joins("left join user on user.id = friend_request.to_id").
		Where("friend_request.from_id = ?", fromId).
		Order("friend_request.created_at desc").
		Find(&friendReqs).Error; err != nil {
		return nil, err
	}
	return friendReqs, nil
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
