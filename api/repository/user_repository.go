package repository

import (
	"api/model/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*entity.Users, error)
	GetFullyAssociativeById(ctx context.Context, id string) (*entity.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) GetById(ctx context.Context, id string) (*entity.Users, error) {
	result := &entity.Users{}
	if err := u.db.WithContext(ctx).Where("user_id = ?", id).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) GetFullyAssociativeById(ctx context.Context, id string) (*entity.Users, error) {
	result := &entity.Users{}
	q := u.db.WithContext(ctx)
	q = q.Preload("Banners")
	q = q.Preload("UserGreetings")
	q = q.Preload("Accounts")
	q = q.Preload("Accounts.AccountBalances")
	q = q.Preload("Accounts.AccountDetails")
	q = q.Preload("Accounts.AccountFlags")
	q = q.Preload("DebitCards")
	q = q.Preload("DebitCards.DebitCardDesign")
	q = q.Preload("DebitCards.DebitCardDetails")
	q = q.Preload("DebitCards.DebitCardStatus")
	q = q.Preload("Transactions")

	if err := q.Where("user_id = ?", id).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
