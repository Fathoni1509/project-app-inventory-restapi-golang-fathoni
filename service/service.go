package service

import "project-app-inventory-restapi-golang-fathoni/repository"

type Service struct {
	CategoryService CategoryService
	WarehouseService WarehouseService
	ShelveService ShelveService
	UserService UserService
}

func NewService(repo repository.Repository) Service {
	return Service{
		CategoryService: NewCategoryService(repo),
		WarehouseService: NewWarehouseService(repo),
		ShelveService: NewShelveService(repo),
		UserService: NewUserService(repo),
	}
}