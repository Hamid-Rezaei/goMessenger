package repository

import (
	"context"
	"errors"
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

func (g *GroupRepository) GetGroupByOwnerID(ctx context.Context, userId uint, groupId uint) (*model.Group, error) {
	var group model.Group

	if err := g.db.Where("ownerId = ? AND id = ?", userId, groupId).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &group, nil
}

func (g *GroupRepository) Delete(ctx context.Context, groupId uint) error {
	if err := g.db.Delete(&model.Group{}, groupId).Error; err != nil {
		return err
	}
	return nil
}
