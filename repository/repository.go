package repository

import (
	"context"

	"github.com/Prthmesh6/bidding_service/models"
)

type Repo interface {
	Add(ctx context.Context, user models.User) (err error)
	Delete(ctx context.Context, userId int64) (deletedUser models.User, err error)
	Update(ctx context.Context, user models.User) (updatedUser models.User, err error)
	GetUsers(ctx context.Context) (users []models.User, err error)
}
