package service

import (
	"context"
	"errors"

	"github.com/Prthmesh6/bidding_service/models"
	"github.com/Prthmesh6/bidding_service/repository"
)

type buyerManager struct {
	buyerRepo repository.Repo
}

func CreatebuyerMgr(buyerRepo repository.Repo) *buyerManager {
	return &buyerManager{
		buyerRepo: buyerRepo,
	}
}

func (b *buyerManager) Add(ctx context.Context, user models.User) (err error) {
	if (user == models.User{}) {
		return errors.New(models.ErrInvalidArguments)
	}

	err = b.buyerRepo.Add(ctx, user)

	return
}

func (b *buyerManager) Delete(ctx context.Context, userId int64) (deletedUser models.User, err error) {
	if userId <= 0 {
		err = errors.New(models.ErrInvalidArguments)
		return
	}

	deletedUser, err = b.buyerRepo.Delete(ctx, userId)
	return
}

func (b *buyerManager) GetUsers(ctx context.Context) (sellers []models.User, err error) {
	sellers, err = b.buyerRepo.GetUsers(ctx)
	return
}
