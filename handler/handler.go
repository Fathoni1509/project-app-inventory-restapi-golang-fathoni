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
	ShelveHandler ShelveHandler
	UserHandler UserHandler
	ProductHandler ProductHandler
}

func NewHandler(service service.Service, config utils.Configuration) Handler {
	return Handler{
		// HandlerAuth: NewAuthHandler(service),
		// HandlerMenu:       NewMenuHandler(),
		CategoryHandler: NewCategoryHandler(service.CategoryService, config),
		WarehouseHandler: NewWarehouseHandler(service.WarehouseService, config),
		ShelveHandler: NewShelveHandler(service.ShelveService, config),
		UserHandler: NewUserHandler(service.UserService, config),
		ProductHandler: NewProductHandler(service.ProductService, config),
	}
}
