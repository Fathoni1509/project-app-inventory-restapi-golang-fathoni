package service

import (
	// "errors"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type InventoryService interface {
	GetListInventorys() ([]dto.InventoryResponse, error)
	GetListInventoryById(inventory_id int) (dto.InventoryResponse, error)
	AddInventory(inventory *model.Inventory) error
	UpdateInventory(inventory_id int, inventory *model.Inventory) error 
	DeleteInventory(inventory_id int) error
}

type inventoryService struct {
	Repo repository.Repository
}

func NewInventoryService(repo repository.Repository) InventoryService {
	return &inventoryService{Repo: repo}
}

// service get list inventorys
func (inv *inventoryService) GetListInventorys() ([]dto.InventoryResponse, error) {
	return inv.Repo.InventoryRepo.GetListInventorys()
}

// servide get list inventory by id
func (inv *inventoryService) GetListInventoryById(inventory_id int) (dto.InventoryResponse, error) {
	return inv.Repo.InventoryRepo.GetListInventoryById(inventory_id)
}

// service add inventory
func (inv *inventoryService) AddInventory(inventory *model.Inventory) error {
	_, err := inv.Repo.ProductRepo.GetListProductById(inventory.ProductId)
	if err != nil {
		return errors.New("product_id is invalid or does not exist")
	}

	_, err = inv.Repo.ShelveRepo.GetListShelveById(inventory.ShelveId)
	if err != nil {
		return errors.New("shelve_id is invalid or does not exist")
	}

	err = inv.Repo.InventoryRepo.AddInventory(inventory)
	if err != nil {
		return err
	}
	
	return nil
}

// service update inventory by ID
func (inv *inventoryService) UpdateInventory(inventory_id int, inventory *model.Inventory) error {
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

	_, err := inv.Repo.ProductRepo.GetListProductById(inventory.ProductId)
	if err != nil {
		return errors.New("product_id is invalid or does not exist")
	}

	_, err = inv.Repo.ShelveRepo.GetListShelveById(inventory.ShelveId)
	if err != nil {
		return errors.New("shelve_id is invalid or does not exist")
	}

	err = inv.Repo.InventoryRepo.UpdateInventory(inventory_id, inventory)
	if err != nil {
		return err
	}

	return nil
}

// service delete inventory by ID
func (inv *inventoryService) DeleteInventory(inventory_id int) error {
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
	return inv.Repo.InventoryRepo.DeleteInventory(inventory_id)
}