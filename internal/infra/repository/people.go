package repository

import (
	"context"
	"errors"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"gorm.io/gorm"
)

type PeopleRepository struct {
	db *gorm.DB
}

func NewPeopleRepo(db *gorm.DB) *PeopleRepository {
	return &PeopleRepository{
		db: db,
	}
}

func (pr *PeopleRepository) Create(ctx context.Context, model model.People) (*model.People, error) {
	tx := pr.db.WithContext(ctx).Begin()

	people := model
	if err := tx.Create(&people).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &people, tx.Commit().Error
}

func (pr *PeopleRepository) Get(ctx context.Context, userId uint, chatId uint) (*model.People, error) {
	var people model.People

	if err := pr.db.Where("user_id = ? and chat_id = ?", userId, chatId).Find(&people).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &people, nil
}

func (pr *PeopleRepository) GetChatUsers(ctx context.Context, chatId uint) ([]uint, error) {
	var people *[]model.People

	if err := pr.db.Where("chat_id = ?", chatId).Find(&people).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var result = []uint{}
	for _, element := range *people {
		result = append(result, element.UserID)
	}
	return result, nil
}
