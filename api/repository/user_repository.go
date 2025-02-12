package repository

import (
	"api/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id string) (*entity.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

// GetById implements UserRepository.
func (u *userRepository) GetById(id string) (*entity.Users, error) {
	result := &entity.Users{}
	if err := u.db.Preload("Banners", "banners.user_id = users.id", id).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
