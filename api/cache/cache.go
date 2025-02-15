package cache

import (
	"errors"
	"sync"
	"time"
)

type LoginCache interface {
	IsExist(userID string) bool
	AddUser(userID string) error
	Increment(userID string) error
	GetCount(userID string) int
	Reset(userID string)
	ResetAfter1Minute(userID string)
	InitCacheReset()
}

type loginCache struct {
	mu       sync.Mutex
	loginMap map[string]int
}

func NewLoginCache() LoginCache {
	return &loginCache{
		loginMap: make(map[string]int),
	}
}

func (lc *loginCache) IsExist(userID string) bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if _, ok := lc.loginMap[userID]; ok {
		return true
	}

	return false
}

func (lc *loginCache) AddUser(userID string) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if _, ok := lc.loginMap[userID]; !ok {
		lc.loginMap[userID] = 0
		return nil
	}

	return errors.New("user already exist")
}

func (lc *loginCache) Increment(userID string) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.loginMap[userID]++
	return nil
}

func (lc *loginCache) GetCount(userID string) int {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	return lc.loginMap[userID]
}

func (lc *loginCache) Reset(userID string) {
	lc.mu.Lock()
	delete(lc.loginMap, userID)
	lc.mu.Unlock()
}

func (lc *loginCache) ResetAfter1Minute(userID string) {
	time.AfterFunc(1*time.Minute, func() {
		lc.mu.Lock()
		delete(lc.loginMap, userID)
		lc.mu.Unlock()
	})
}

func (lc *loginCache) InitCacheReset() {
	go lc.resetDaily()
}

func (lc *loginCache) resetDaily() {
	for {
		time.Sleep(24 * time.Hour)
		lc.mu.Lock()
		lc.loginMap = make(map[string]int)
		lc.mu.Unlock()
	}
}
