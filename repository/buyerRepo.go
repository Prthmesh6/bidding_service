package repository

import (
	"context"
	"sync"

	"github.com/Prthmesh6/bidding_service/models"
)

type buyerRepo struct {
	storageMap map[int64]models.User
	mutex      sync.RWMutex
}

func CreateBuyerRepo() *buyerRepo {
	return &buyerRepo{
		storageMap: make(map[int64]models.User),
		mutex:      sync.RWMutex{},
	}
}

func (b *buyerRepo) Add(ctx context.Context, user models.User) (err error) {
	b.mutex.Lock()
	b.storageMap[user.Id] = user
	b.mutex.Unlock()
	return
}

func (b *buyerRepo) Delete(ctx context.Context, userId int64) (deletedUser models.User, err error) {
	b.mutex.Lock()
	delete(b.storageMap, userId)
	b.mutex.Unlock()
	return
}

func (b *buyerRepo) Update(ctx context.Context, user models.User) (updatedUser models.User, err error) {
	b.mutex.Lock()
	b.storageMap[user.Id] = user
	b.mutex.Unlock()
	updatedUser = user
	return
}

func (b *buyerRepo) GetUsers(ctx context.Context) (users []models.User, err error) {
	b.mutex.RLock()
	for _, val := range b.storageMap {
		users = append(users, val)
	}
	b.mutex.RUnlock()

	return
}
