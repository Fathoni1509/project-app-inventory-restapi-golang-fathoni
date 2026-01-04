package service

import (
	// "errors"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type WarehouseService interface {
	GetListWarehouses() ([]model.Warehouse, error)
	GetListWarehouseById(warehouse_id int) (model.Warehouse, error)
	AddWarehouse(warehouse *model.Warehouse) error
	UpdateWarehouse(warehouse_id int, warehouse *model.Warehouse) error 
	DeleteWarehouse(warehouse_id int) error
}

type warehouseService struct {
	Repo repository.Repository
}

func NewWarehouseService(repo repository.Repository) WarehouseService {
	return &warehouseService{Repo: repo}
}

// service get list warehouses
func (war *warehouseService) GetListWarehouses() ([]model.Warehouse, error) {
	return war.Repo.WarehouseRepo.GetListWarehouses()
}

// servide get list warehouse by id
func (war *warehouseService) GetListWarehouseById(warehouse_id int) (model.Warehouse, error) {
	return war.Repo.WarehouseRepo.GetListWarehouseById(warehouse_id)
}

// service add warehouse
func (war *warehouseService) AddWarehouse(warehouse *model.Warehouse) error {
	return war.Repo.WarehouseRepo.AddWarehouse(warehouse)
}

// service update category by ID
func (war *warehouseService) UpdateWarehouse(warehouse_id int, warehouse *model.Warehouse) error {
	// categories, err := cat.GetListCategories()
	// if err != nil {
	// 	return err
	// }
	// for _, c := range categories {
	// 	if category_id == c.CategoryId {
	// 		return cat.Repo.CategoryRepo.UpdateCategory(category_id, category)
	// 	}
	// }

	// return errors.New("id category not found")

	return war.Repo.WarehouseRepo.UpdateWarehouse(warehouse_id, warehouse)
}

// service delete category by ID
func (war *warehouseService) DeleteWarehouse(warehouse_id int) error {
	// categories, err := cat.Repo.CategoryRepo.GetListCategories()
	// if err != nil {
	// 	return err
	// }
	// for _, c := range categories {
	// 	if category_id == c.CategoryId {
	// 		return cat.Repo.CategoryRepo.DeleteCategory(category_id, category)
	// 	}
	// }

	// return errors.New("id category not found")
	return war.Repo.WarehouseRepo.DeleteWarehouse(warehouse_id)
}