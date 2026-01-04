package service

import (
	// "errors"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type ShelveService interface {
	GetListShelves() ([]model.Shelve, error)
	GetListShelveById(shelve_id int) (model.Shelve, error)
	AddShelve(shelve *model.Shelve) error
	UpdateShelve(shelve_id int, shelve *model.Shelve) error 
	DeleteShelve(shelve_id int) error
}

type shelveService struct {
	Repo repository.Repository
}

func NewShelveService(repo repository.Repository) ShelveService {
	return &shelveService{Repo: repo}
}

// service get list shelves
func (sh *shelveService) GetListShelves() ([]model.Shelve, error) {
	return sh.Repo.ShelveRepo.GetListShelves()
}

// servide get list shelve by id
func (sh *shelveService) GetListShelveById(shelve_id int) (model.Shelve, error) {
	return sh.Repo.ShelveRepo.GetListShelveById(shelve_id)
}

// service add shelve
func (sh *shelveService) AddShelve(shelve *model.Shelve) error {
	_, err := sh.Repo.WarehouseRepo.GetListWarehouseById(shelve.WarehouseId)
	if err != nil {
		return errors.New("warehouse_id is invalid or does not exist")
	}

	err = sh.Repo.ShelveRepo.AddShelve(shelve)
	if err != nil {
		return err
	}
	
	return nil
}

// service update shelve by ID
func (sh *shelveService) UpdateShelve(shelve_id int, shelve *model.Shelve) error {
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

	_, err := sh.Repo.WarehouseRepo.GetListWarehouseById(shelve.WarehouseId)
	if err != nil {
		return errors.New("warehous_id is invalid or does not exist")
	}

	err = sh.Repo.ShelveRepo.UpdateShelve(shelve_id, shelve)
	if err != nil {
		return err
	}

	return nil
}

// service delete shelve by ID
func (sh *shelveService) DeleteShelve(shelve_id int) error {
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
	return sh.Repo.ShelveRepo.DeleteShelve(shelve_id)
}