package repo

import (
	"github.com/google/uuid"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	"gorm.io/gorm"
	"time"
)

type RefreshRepo interface {
	CreateRefreshToken(userId string) (string, error)
}

type refreshRepo struct {
	db *gorm.DB
}

func NewRefreshRepo() RefreshRepo {
	return &refreshRepo{
		db: db.DB,
	}
}

// CreateRefreshToken generate refresh token and save to database
func (r *refreshRepo) CreateRefreshToken(userId string) (string, error) {
	// 生成一个新的 refresh token
	token := &model.UserRefreshToken{
		ID:           uuid.NewString(),
		UserID:       userId,
		RefreshToken: uuid.NewString(),
		ExpiresAt:    time.Now().Add(30 * 24 * time.Hour).Unix(), // 30 天后过期
		Revoked:      true,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	if err := r.db.Create(token); err != nil {
		return "", err.Error
	}
	return token.RefreshToken, nil
}

// ValidateToken validate refresh token
func (r *refreshRepo) ValidateToken(userId string, refreshToken string) bool {
	var count int64
	r.db.Model(&model.UserRefreshToken{}).
		Where("user_id = ? and refresh_token = ? and revoked = false AND expires_at > ?", userId, refreshToken, time.Now().Unix()).Count(&count)
	return count > 0
}
