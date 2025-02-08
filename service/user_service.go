package service

import (
	"errors"

	user "github.com/adosalkanovicc/go_crud/model"
	"github.com/adosalkanovicc/go_crud/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (us *UserService) GetUsers(page, limit int) ([]*user.User, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("invalid page or limit")
	}
	offset := (page - 1) * limit
	return us.UserRepo.GetUsers(limit, offset)
}

func (us *UserService) CreateUser(u *user.User) error {
	return us.UserRepo.CreateUser(u)
}

func (us *UserService) UpdateUser(u *user.User) error {
	if u.Id <= 0 {
		return errors.New("invalid ID")
	}
	return us.UserRepo.UpdateUser(u)
}

func (us *UserService) DeleteUser(id int) error {
	return us.UserRepo.DeleteUser(id)
}
