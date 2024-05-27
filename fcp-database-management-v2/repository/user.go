package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	result := u.db.First(&user, "Username = ? AND Password = ?", user.Username, user.Password)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
