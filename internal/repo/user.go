package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/pkg/db"
	utils2 "go-chat-server/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByMail(ctx context.Context, email string) (*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)
	CreateUserByMail(ctx context.Context, email string, firstPwd string) (*model.User, error)
	GetByGithubId(githubId int64) (*model.User, error)
	CreateGithubUser(githubUser map[string]interface{}) (*model.User, error)
	GetUserByKeyword(ctx context.Context, keyword string, fromId string) (*v1.AddUserResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db.DB}
}

func (r *userRepository) GetByMail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.WithContext(ctx).Where("mail = ?", email).First(user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by mail: %v", err)
	}
	return user, nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by mail: %v", err)
	}
	return user, nil
}

func (r *userRepository) CreateUserByMail(ctx context.Context, email string, firstPwd string) (*model.User, error) {
	user := &model.User{
		ID:       uuid.NewString(),
		Mail:     email,
		Name:     email,
		NickName: utils2.GenerateUsername(8),
	}
	if firstPwd != "" {
		user.Password = utils2.ToHash(firstPwd)
	} else {
		user.Password = utils2.ToHash("123456")
	}
	if err := r.db.WithContext(ctx).Omit("Phone").Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user by email: %v", err)
	}
	return user, nil
}

func (r *userRepository) GetByGithubId(githubId int64) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("github_id = ?", githubId).First(user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by github id: %v", err)
	}
	return user, nil
}

func (r *userRepository) CreateGithubUser(githubUser map[string]interface{}) (*model.User, error) {
	user := &model.User{
		ID:       uuid.NewString(),
		Mail:     githubUser["email"].(string),
		Name:     githubUser["email"].(string),
		Avatar:   githubUser["avatar_url"].(string),
		NickName: githubUser["name"].(string),
		GithubId: int64(githubUser["id"].(float64)),
	}
	user.Password = utils2.ToHash("123456")
	if err := r.db.Omit("Phone").Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user by email: %v", err)
	}
	return user, nil
}

func (r *userRepository) GetUserByKeyword(ctx context.Context, keyword string, fromId string) (*v1.AddUserResponse, error) {
	var user v1.AddUserResponse
	err := r.db.WithContext(ctx).
		Table("user").Select("id", "nick_name", "mail", "name", "avatar", "phone"). // 👈 只查这几个列
		Where("nick_name = ? OR mail = ?", keyword, keyword).
		Limit(1).
		Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// 查询是否已经好友申请过及申请状态
	var friendRequest model.FriendRequest
	err1 := r.db.WithContext(ctx).Table("friend_request").Select("status", "created_at").
		Where("from_id = ? AND to_id = ?", fromId, user.ID).
		Order("created_at desc").
		Limit(1).Find(&friendRequest).Error
	if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
		friendRequest.Status = "not-applied"
	}
	if friendRequest.Status == "rejected" && (utils2.GetNowTimeUnix()-friendRequest.CreatedAt) > 24*60*60*1000 {
		friendRequest.Status = "not-applied"
	}
	user.Status = friendRequest.Status
	return &user, nil
}
