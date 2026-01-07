package handler

import (
	"project-app-inventory-restapi-golang-fathoni/service"
	"project-app-inventory-restapi-golang-fathoni/utils"
)

type Handler struct {
	CategoryHandler  CategoryHandler
	WarehouseHandler WarehouseHandler
	ShelveHandler    ShelveHandler
	UserHandler      UserHandler
	ProductHandler   ProductHandler
	SaleHandler      SaleHandler
	ReportHandler    ReportHandler
}

func NewHandler(service service.Service, config utils.Configuration) Handler {
	return Handler{
		CategoryHandler:  NewCategoryHandler(service.CategoryService, config),
		WarehouseHandler: NewWarehouseHandler(service.WarehouseService, config),
		ShelveHandler:    NewShelveHandler(service.ShelveService, config),
		UserHandler:      NewUserHandler(service.UserService, config),
		ProductHandler:   NewProductHandler(service.ProductService, config),
		SaleHandler:      NewSaleHandler(service.SaleService, config),
		ReportHandler:    NewReportHandler(service.ReportService, config),
	}
}
