package service

import (
	"context"
	"errors"
	"go-chat-server/internal/repo"
	"go-chat-server/internal/utils"
	"gorm.io/gorm"
)

type UserService interface {
	GenerateToken(ctx context.Context, email string) (string, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func (u *userService) GenerateToken(ctx context.Context, email string) (string, error) {
	user, err := u.userRepo.GetByMail(ctx, email)
	if err != nil {
		if errors.As(err, gorm.ErrRecordNotFound) {
			// User not found, create a new user
			user, err = u.userRepo.CreateUserByMail(ctx, email)
			if err != nil {
				return "", err
			}
		} else {
			return "", err

		}
	}
	// Generate a token for the user
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
