package service

import (
	"context"
	"errors"

	"github.com/Prthmesh6/bidding_service/models"
	"github.com/Prthmesh6/bidding_service/repository"
)

type SellerManager struct {
	sellerRepo repository.Repo
}

func CreateSellerMgr(SellerRepo repository.Repo) *SellerManager {
	return &SellerManager{
		sellerRepo: SellerRepo,
	}
}

func (s *SellerManager) Add(ctx context.Context, user models.User) (err error) {
	if (user == models.User{}) {
		return errors.New(models.ErrInvalidArguments)
	}

	err = s.sellerRepo.Add(ctx, user)

	return
}

func (s *SellerManager) Delete(ctx context.Context, userId int32) (deletedUser models.User, err error) {
	if userId <= 0 {
		err = errors.New(models.ErrInvalidArguments)
		return
	}

	deletedUser, err = s.sellerRepo.Delete(ctx, userId)
	return
}

func (s *SellerManager) Update(ctx context.Context, user models.User) (updatedUser models.User, err error) {
	updatedUser, err = s.sellerRepo.Update(ctx, user)
	return
}

func (s *SellerManager) GetUsers(ctx context.Context) (sellers []models.User, err error) {
	sellers, err = s.sellerRepo.GetUsers(ctx)
	return
}
