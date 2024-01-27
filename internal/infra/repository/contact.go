package repository

import (
	"context"
	"errors"

	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepo(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		db: db,
	}
}

func (cr *ContactRepository) Create(ctx context.Context, model model.Contact) (*model.Contact, error) {
	tx := cr.db.WithContext(ctx).Begin()

	contact := model
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &contact, tx.Commit().Error
}

func (cr *ContactRepository) GetById(_ context.Context, user_id uint, contact_id uint) (*model.Contact, error) {
	var contact = model.Contact{UserId: user_id, ContactId: contact_id}

	if err := cr.db.First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &contact, nil
}

func (cr *ContactRepository) GetList(ctx context.Context, user_id uint) (*[]model.Contact, error) {
	var contacts []model.Contact

	if err := cr.db.Where("userid = ?", user_id).Find(&contacts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &contacts, nil
}

func (cr *ContactRepository) Delete(ctx context.Context, user_id uint, contact_id uint) error {
	tx := cr.db.WithContext(ctx).Begin()

	if err := tx.Delete(&model.Contact{UserId: user_id, ContactId: contact_id}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
