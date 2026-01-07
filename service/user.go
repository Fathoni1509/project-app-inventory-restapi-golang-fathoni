package service

import (
	// "errors"
	// "errors"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type UserService interface {
	GetListUsers() ([]model.User, error)
	GetListUserById(user_id int) (model.User, error)
	AddUser(user *model.User) error
	UpdateUser(user_id int, user *model.User) error
	DeleteUser(user_id int) error
}

type userService struct {
	Repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{Repo: repo}
}

// service get list users
func (us *userService) GetListUsers() ([]model.User, error) {
	return us.Repo.UserRepo.GetListUsers()
}

// servide get list user by id
func (us *userService) GetListUserById(user_id int) (model.User, error) {
	return us.Repo.UserRepo.GetListUserById(user_id)
}

// service add user
func (us *userService) AddUser(user *model.User) error {
	return us.Repo.UserRepo.AddUser(user)
}

// service update user by ID
func (us *userService) UpdateUser(user_id int, user *model.User) error {
	return us.Repo.UserRepo.UpdateUser(user_id, user)
}

// service delete user by ID
func (us *userService) DeleteUser(user_id int) error {
	return us.Repo.UserRepo.DeleteUser(user_id)
}
