package repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-chat-server/internal/model"
	"go-chat-server/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByMail(ctx context.Context, email string) (*model.User, error)
	CreateUserByMail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByMail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.WithContext(ctx).Where("mail = ?", email).First(user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by mail: %v", err)
	}
	return user, nil
}

func (r *userRepository) CreateUserByMail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{
		ID:       uuid.NewString(),
		Mail:     email,
		NickName: utils.GenerateUsername(8),
		Password: utils.ToHash("123456"),
	}
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user by email: %v", err)
	}
	return user, nil
}
