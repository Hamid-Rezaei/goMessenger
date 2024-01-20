package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, model model.User) (uint, error) {
	tx := ur.db.WithContext(ctx).Begin()

	user := model
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return user.ID, tx.Commit().Error
}
