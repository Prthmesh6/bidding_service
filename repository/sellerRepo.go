package repository

import (
	"context"
	"sync"

	"github.com/Prthmesh6/bidding_service/models"
)

type sellerRepo struct {
	storageMap map[int64]models.User
	mutex      sync.RWMutex
}

func CreateSellerRepo() *sellerRepo {
	return &sellerRepo{
		storageMap: make(map[int64]models.User),
		mutex:      sync.RWMutex{},
	}
}

func (s *sellerRepo) Add(ctx context.Context, user models.User) (err error) {
	s.mutex.Lock()
	s.storageMap[user.Id] = user
	s.mutex.Unlock()
	return
}

func (s *sellerRepo) Delete(ctx context.Context, userId int64) (deletedUser models.User, err error) {
	s.mutex.Lock()
	delete(s.storageMap, userId)
	s.mutex.Unlock()
	return
}

func (s *sellerRepo) Update(ctx context.Context, user models.User) (updatedUser models.User, err error) {
	s.mutex.Lock()
	s.storageMap[user.Id] = user
	s.mutex.Unlock()
	updatedUser = user
	return
}

func (s *sellerRepo) GetUsers(ctx context.Context) (users []models.User, err error) {
	s.mutex.RLock()
	for _, val := range s.storageMap {
		users = append(users, val)
	}
	s.mutex.RUnlock()
	return
}
