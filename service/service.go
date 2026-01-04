package service

import "project-app-inventory-restapi-golang-fathoni/repository"

type Service struct {
	CategoryService CategoryService
	WarehouseService WarehouseService
}

func NewService(repo repository.Repository) Service {
	return Service{
		CategoryService: NewCategoryService(repo),
		WarehouseService: NewWarehouseService(repo),
	}
}