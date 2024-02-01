package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (g *GroupRepository) Create(ctx context.Context, model model.Group) (*model.Group, error) {
	tx := g.db.WithContext(ctx).Begin()

	group := model
	if err := tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &group, tx.Commit().Error
}
