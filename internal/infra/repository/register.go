package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, model model.User) (uint, error)
	GetByUsernamePhone(_ context.Context, username string, phone string) (*model.User, error)
	GetUserByID(_ context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, user *model.User, id uint) error
}
