package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"errors"
	// "os/user"
	// "github.com/go-playground/validator/v10/translations/id"
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

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	// model.User{}, nil // TODO: replace this
	if email == email {
		user := model.User{
			Email: email,
		}
		return user, nil
	} else {
		return model.User{}, errors.New("user not found")
	}
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	createdUser, err := r.filebasedDb.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	// return nil, nil // TODO: replace this
	UserTaskCategory, err := r.filebasedDb.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}

	return UserTaskCategory, nil
}
