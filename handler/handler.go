package handler

import (
	"project-app-inventory-restapi-golang-fathoni/service"
	"project-app-inventory-restapi-golang-fathoni/utils"
)

type Handler struct {
	// HandlerAuth       AuthHandler
	// HandlerMenu       MenuHandler
	CategoryHandler CategoryHandler
	WarehouseHandler WarehouseHandler
}

func NewHandler(service service.Service, config utils.Configuration) Handler {
	return Handler{
		// HandlerAuth: NewAuthHandler(service),
		// HandlerMenu:       NewMenuHandler(),
		CategoryHandler: NewCategoryHandler(service.CategoryService, config),
		WarehouseHandler: NewWarehouseHandler(service.WarehouseService, config),
	}
}
