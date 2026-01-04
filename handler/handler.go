package handler

import (
	"project-app-inventory-restapi-golang-fathoni/service"
	"project-app-inventory-restapi-golang-fathoni/utils"
)

type Handler struct {
	// HandlerAuth       AuthHandler
	// HandlerMenu       MenuHandler
	CategoryHandler CategoryHandler
}

func NewHandler(service service.Service, config utils.Configuration) Handler {
	return Handler{
		// HandlerAuth: NewAuthHandler(service),
		// HandlerMenu:       NewMenuHandler(),
		CategoryHandler: NewCategoryHandler(service.CategoryService, config),
	}
}
