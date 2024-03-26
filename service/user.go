package service

import (
	"context"

	"github.com/Prthmesh6/bidding_service/models"
)

type User interface {
	Add(ctx context.Context, user models.User) (err error)
	Delete(ctx context.Context, userId int32) (deletedUser models.User, err error)
	Update(ctx context.Context, user models.User) (updatedUser models.User, err error)
	GetUsers(ctx context.Context) (userNames []models.User, err error)
}
