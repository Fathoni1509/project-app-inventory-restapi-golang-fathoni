package service

import "project-app-inventory-restapi-golang-fathoni/repository"

type Service struct {
	CategoryService CategoryService
}

func NewService(repo repository.Repository) Service {
	return Service{
		CategoryService: NewCategoryService(repo),
	}
}