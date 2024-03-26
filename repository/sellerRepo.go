package repository

import (
	"context"

	"github.com/Prthmesh6/bidding_service/models"
)

type sellerRepo struct {
	storageMap map[int32]models.User
}

func CreateSellerRepo() *sellerRepo {
	return &sellerRepo{
		storageMap: make(map[int32]models.User),
	}
}

func (s *sellerRepo) Add(ctx context.Context, user models.User) (err error) {
	s.storageMap[user.Id] = user
	return
}

func (s *sellerRepo) Delete(ctx context.Context, userId int32) (deletedUser models.User, err error) {
	delete(s.storageMap, userId)
	return
}

func (s *sellerRepo) Update(ctx context.Context, user models.User) (updatedUser models.User, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *sellerRepo) GetUsers(ctx context.Context) (users []models.User, err error) {
	for _, val := range s.storageMap {
		users = append(users, val)
	}

	return
}
