package service

import (
	// "errors"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
)

type CategoryService interface {
	GetListCategories() ([]model.Category, error)
	GetListCategoryById(category_id int) (model.Category, error)
	AddCategory(category *model.Category) error
	UpdateCategory(category_id int, category *model.Category) error 
	DeleteCategory(category_id int) error
}

type categoryService struct {
	Repo repository.Repository
}

func NewCategoryService(repo repository.Repository) CategoryService {
	return &categoryService{Repo: repo}
}

// service get list categories
func (cat *categoryService) GetListCategories() ([]model.Category, error) {
	return cat.Repo.CategoryRepo.GetListCategories()
}

// servide get list category by id
func (cat *categoryService) GetListCategoryById(category_id int) (model.Category, error) {
	return cat.Repo.CategoryRepo.GetListCategoryById(category_id)
}

// service add category
func (cat *categoryService) AddCategory(category *model.Category) error {
	return cat.Repo.CategoryRepo.AddCategory(category)
}

// service update category by ID
func (cat *categoryService) UpdateCategory(category_id int, category *model.Category) error {
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

	return cat.Repo.CategoryRepo.UpdateCategory(category_id, category)
}

// service delete category by ID
func (cat *categoryService) DeleteCategory(category_id int) error {
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
	return cat.Repo.CategoryRepo.DeleteCategory(category_id)
}