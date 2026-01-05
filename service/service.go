package service

import "project-app-inventory-restapi-golang-fathoni/repository"

type Service struct {
	CategoryService CategoryService
	WarehouseService WarehouseService
	ShelveService ShelveService
	UserService UserService
	ProductService ProductService
}

func NewService(repo repository.Repository) Service {
	return Service{
		CategoryService: NewCategoryService(repo),
		WarehouseService: NewWarehouseService(repo),
		ShelveService: NewShelveService(repo),
		UserService: NewUserService(repo),
		ProductService: NewProductService(repo),
	}
}