package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"errors"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	filebasedDb *filebased.Data
}

func NewUserRepo(filebasedDb *filebased.Data) *userRepository {
	return &userRepository{filebasedDb}
}

// GetUserByEmail retrieves a user by their email
func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	user, err := r.filebasedDb.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	if (user == model.User{}) {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

// CreateUser creates a new user
func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	createdUser, err := r.filebasedDb.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	return createdUser, nil
}

// GetUserTaskCategory retrieves all user task categories
func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	taskCategories, err := r.filebasedDb.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}
	return taskCategories, nil
}
