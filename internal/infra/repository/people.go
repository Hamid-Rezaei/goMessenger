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

func (pr *PeopleRepository) SetNewMessageToZero(ctx context.Context, chatId uint, userId uint) error {
	pr.db.Raw("UPDATE Peoples WHERE chat_id=? and user_id = ? SET new_messages=0", chatId, userId)
	return nil
}

func (pr *PeopleRepository) AddNewMessages(ctx context.Context, chatId uint, userId uint) error {
	res, err := pr.GetChatUsers(ctx, chatId)
	if err != nil {
		return err
	}
	var user_id = 0
	for _, element := range res {
		if element != userId {
			user_id = int(element)
			break
		}
	}
	var people model.People
	err1 := pr.db.Where("user_id = ? and chat_id = ?", user_id, chatId).First(&people).Error
	if err1 != nil {
		return err1
	}
	pr.db.Raw("UPDATE Peoples WHERE chat_id=? and user_id = ? SET new_messages=?", chatId, user_id, people.NewMessages+1)
	return nil
}

func (pr *PeopleRepository) GetNewMessagesCount(_ context.Context, chatId uint, userId uint) (int, error) {
	var newMessagesCount int
	if err := pr.db.Model(&model.People{}).
		Where("chat_id = ? AND user_id = ?", chatId, userId).
		Pluck("new_messages", &newMessagesCount).
		Error; err != nil {
		return 0, err
	}

	return newMessagesCount, nil
}
