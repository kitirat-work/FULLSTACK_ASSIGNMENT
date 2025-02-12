package service

import (
	"api/model/entity"
	"api/repository"
)

type UserService interface {
	GetById(id string) (*entity.Users, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// GetById implements UserService.
func (u *userService) GetById(id string) (*entity.Users, error) {
	return u.userRepo.GetById(id)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
