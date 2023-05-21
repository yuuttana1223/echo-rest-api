package repository

import (
	"echo-rest-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r UserRepositoryImpl) GetUserByEmail(user *model.User, email string) error {
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (r UserRepositoryImpl) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
