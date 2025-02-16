package service

import (
	"api/cache"
	"api/model/entity"
	"api/repository"
	"context"
	"errors"
	"fmt"
)

type AuthService interface {
	LoginPin(ctx context.Context, id string, pin string) (*entity.Users, error)
}

type authService struct {
	cache    cache.LoginCache
	userRepo repository.UserRepository
	limit    int
}

func NewAuthService(cache cache.LoginCache, userRepo repository.UserRepository) AuthService {
	return &authService{
		cache:    cache,
		userRepo: userRepo,
		limit:    3,
	}
}

func (a *authService) LoginPin(ctx context.Context, id string, pin string) (*entity.Users, error) {
	if !a.cache.IsExist(id) {
		if err := a.cache.AddUser(id); err != nil {
			return nil, err
		}
	}

	if a.cache.GetCount(id) >= a.limit {
		a.cache.ResetAfter1Minute(id)
		return nil, errors.New("Account locked. Try again 1 minute later.")
	}

	// Hardcoded PIN for testing purposes.
	if pin != "000000" {
		if err := a.cache.Increment(id); err != nil {
			return nil, err
		}

		attemptsLeft := a.limit - a.cache.GetCount(id)
		if attemptsLeft == 0 {
			a.cache.ResetAfter1Minute(id)
			return nil, errors.New("Account locked. Try again 1 minute later.")
		}
		return nil, errors.New(fmt.Sprintf("Invalid PIN. You have %d attempts left.", attemptsLeft))
	}

	a.cache.Reset(id)

	user, err := a.userRepo.GetFullyAssociativeById(ctx, id)
	if err != nil {
		return nil, err
	}
	// ปกติเราจะไม่ควร return entity ออกไปเลย แต่ในกรณีนี้เราจะให้ผ่านไปก่อน
	// ปกติเราควรสร้าง DTO มาเพื่อให้เป็นตัวกลางระหว่าง entity กับ presentation layer
	// และให้ entity อยู่ใน domain layer เท่านั้น
	// ปกติควรส่งข้อมูล user ไปทาง cookie หรือ session และไม่ควรส่ง entity ไปทาง response
	// แต่ในกรณีนี้เราจะให้ผ่านไปก่อน
	return user, nil
}
