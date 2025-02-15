package service

import (
	"api/model/entity"
	"api/repository"
	"context"
)

type UserService interface {
	GetById(ctx context.Context, id string) (*entity.Users, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// No need to test this function because it's just a wrapper.
func (u *userService) GetById(ctx context.Context, id string) (*entity.Users, error) {
	return u.userRepo.GetById(ctx, id)
}
