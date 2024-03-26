package repository

import (
	"context"

	"github.com/Prthmesh6/bidding_service/models"
)

type buyerRepo struct {
	storageMap map[int32]models.User
}

func CreateBuyerRepo() *buyerRepo {
	return &buyerRepo{
		storageMap: make(map[int32]models.User),
	}
}

func (b *buyerRepo) Add(ctx context.Context, user models.User) (err error) {
	b.storageMap[user.Id] = user
	return
}

func (b *buyerRepo) Delete(ctx context.Context, userId int32) (deletedUser models.User, err error) {
	delete(b.storageMap, userId)
	return
}

func (b *buyerRepo) Update(ctx context.Context, user models.User) (updatedUser models.User, err error) {
	panic("not implemented") // TODO: Implement
}

func (b *buyerRepo) GetUsers(ctx context.Context) (users []models.User, err error) {
	for _, val := range b.storageMap {
		users = append(users, val)
	}

	return
}
