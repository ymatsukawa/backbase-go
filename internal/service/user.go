package service

import (
	"context"

	"github.com/ymatsukawa/backbase-go/internal/entity"
	"github.com/ymatsukawa/backbase-go/internal/repository"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return u.userRepo.GetAllUsers(ctx)
}
