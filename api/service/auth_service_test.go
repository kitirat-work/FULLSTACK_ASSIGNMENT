package service_test

import (
	"api/cache/mocks"
	"api/model/entity"
	mockRepo "api/repository/mocks"
	"api/service"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_LoginPin(t *testing.T) {
	var (
		svc      service.AuthService
		cache    mocks.LoginCache
		userRepo mockRepo.UserRepository

		ctx          context.Context
		resIsExist   bool
		errAddUser   error
		resGetCount  int
		errIncrement error
		resGetUser   *entity.Users
		errGetUser   error
	)

	beforeEach := func() {
		cache = mocks.LoginCache{}
		userRepo = mockRepo.UserRepository{}
		svc = service.NewAuthService(&cache, &userRepo)

		ctx = context.Background()
		resIsExist = false
		errAddUser = nil
		resGetCount = 0
		errIncrement = nil
		resGetUser = nil
		errGetUser = nil

		cache.On("IsExist", mock.Anything).Return(
			func(string) bool {
				return resIsExist
			},
		)

		cache.On("AddUser", mock.Anything).Return(
			func(string) error {
				return errAddUser
			},
		)

		cache.On("GetCount", mock.Anything).Return(
			func(string) int {
				return resGetCount
			},
		)

		cache.On("Increment", mock.Anything).Return(
			func(string) error {
				return errIncrement
			},
		)

		cache.On("Reset", mock.Anything)

		userRepo.On("GetFullyAssociativeById", mock.Anything, mock.Anything).Return(
			func(context.Context, string) *entity.Users {
				return resGetUser
			},
			func(context.Context, string) error {
				return errGetUser
			},
		)

	}

	t.Run("should add user to cache if not exist before", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = false
		errAddUser = nil

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Nil(t, err)
		cache.AssertCalled(t, "AddUser", id)
	})

	t.Run("should not call add user to cache if user already exist", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Nil(t, err)
		cache.AssertNotCalled(t, "AddUser", id)
	})

	t.Run("should return error if add user to cache failed", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = false
		errAddUser = errors.New("error add user")

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Error(t, err)
		cache.AssertCalled(t, "AddUser", id)
	})

	t.Run("should return error when count of login attempt exceed limit", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true
		resGetCount = 3

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Error(t, err)
		cache.AssertCalled(t, "GetCount", id)
	})

	t.Run("should increment count of login attempt", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true
		resGetCount = 2
		errIncrement = nil

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Nil(t, err)
		cache.AssertCalled(t, "Increment", id)
	})

	t.Run("should return error when increment count of login attempt failed", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true
		resGetCount = 2
		errIncrement = errors.New("error increment")

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Error(t, err)
		cache.AssertCalled(t, "Increment", id)
	})

	t.Run("should return error when password is invalid", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "123456"
		resIsExist = true
		resGetCount = 2
		errIncrement = nil

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Error(t, err)
	})

	t.Run("should reset count of login attempt when password is valid", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true
		resGetCount = 2
		errIncrement = nil

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Nil(t, err)
		cache.AssertCalled(t, "Reset", id)
	})

	t.Run("should return user entity when login success", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true
		resGetCount = 2
		errIncrement = nil
		resGetUser = &entity.Users{
			UserID: id,
		}

		user, err := svc.LoginPin(ctx, id, pin)

		assert.Nil(t, err)
		assert.Equal(t, id, user.UserID)
	})

	t.Run("should return error when get user failed", func(t *testing.T) {
		beforeEach()
		id := "123456"
		pin := "000000"
		resIsExist = true
		resGetCount = 2
		errIncrement = nil
		resGetUser = nil
		errGetUser = errors.New("error get user")

		_, err := svc.LoginPin(ctx, id, pin)

		assert.Error(t, err)
	})
}
