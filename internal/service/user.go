package service

import (
	"context"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/internal/repo"
	"go-chat-server/pkg/jwt"
	"go-chat-server/pkg/utils"
)

type UserService interface {
	IsNewUser(ctx context.Context, mail string) bool
	RegisterNewUser(ctx context.Context, request *v1.RegisterByCodeRequest) (model.User, error)
	VerifyPwdWithToken(ctx context.Context, request *v1.LoginByPwdRequest) string
	GetAuthUserProfile(ctx context.Context, userId string) *v1.AuthUserResponse
	GetUserByKeyword(ctx context.Context, keyword string, formId string) (*v1.AddUserResponse, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService() UserService {
	return &userService{repo.NewUserRepository()}
}

func (u *userService) IsNewUser(ctx context.Context, mail string) bool {
	user, _ := u.userRepo.GetByMail(ctx, mail)
	if user != nil {
		return false
	}
	return true
}

func (u *userService) RegisterNewUser(ctx context.Context, request *v1.RegisterByCodeRequest) (model.User, error) {
	user, err := u.userRepo.CreateUserByMail(ctx, request.Mail, request.Password)
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}

func (u *userService) VerifyPwdWithToken(ctx context.Context, request *v1.LoginByPwdRequest) string {
	user, _ := u.userRepo.GetByMail(ctx, request.Mail)
	if user == nil {
		return ""
	}
	if user.Password == utils.ToHash(request.Password) {
		token, err := jwt.GenerateToken(user.ID)
		if err != nil {
			return ""
		}
		return token
	}
	return ""
}

func (u *userService) GetAuthUserProfile(ctx context.Context, userId string) *v1.AuthUserResponse {
	user, err := u.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil
	}
	return &v1.AuthUserResponse{
		ID:       user.ID,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Mail:     user.Mail,
		Phone:    user.Phone,
	}
}

func (u *userService) GetUserByKeyword(ctx context.Context, keyword string, formId string) (*v1.AddUserResponse, error) {
	return u.userRepo.GetUserByKeyword(ctx, keyword, formId)
}
