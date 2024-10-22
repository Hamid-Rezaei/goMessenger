package repository

import (
	"context"
	"errors"
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

func (ur *UserRepository) GetByUsername(_ context.Context, username string) (*model.User, error) {
	var u model.User
	if err := ur.db.Where(&model.User{Username: username}).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (ur *UserRepository) GetUserByID(_ context.Context, id uint) (*model.User, error) {
	var u model.User

	if err := ur.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (ur *UserRepository) GetUsersByID(_ context.Context, ids []uint) ([]model.User, error) {
	var users []model.User

	if err := ur.db.Find(&users, ids).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *model.User, id uint) error {
	tx := ur.db.WithContext(ctx).Begin()

	if err := tx.Model(&model.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	tx := ur.db.WithContext(ctx).Begin()

	if err := tx.Delete(&model.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (ur *UserRepository) SearchUser(_ context.Context, keyword string) (*model.User, error) {
	var u model.User
	q := "%" + keyword + "%"
	if err := ur.db.Where("lower(username) LIKE ?", q).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
