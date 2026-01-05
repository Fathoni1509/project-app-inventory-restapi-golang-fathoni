package service

import (
	// "errors"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
	"project-app-inventory-restapi-golang-fathoni/utils"
)

type ProductService interface {
	GetListProducts(page, limit int) (*[]model.Product, *dto.Pagination, error)
	GetListProductById(product_id int) (model.Product, error)
	AddProduct(product *model.Product) error
	UpdateProduct(product_id int, product *model.Product) error 
	DeleteProduct(product_id int) error
}

type productService struct {
	Repo repository.Repository
}

func NewProductService(repo repository.Repository) ProductService {
	return &productService{Repo: repo}
}

// service get list products
func (pr *productService) GetListProducts(page, limit int) (*[]model.Product, *dto.Pagination, error) {
	products, total, err := pr.Repo.ProductRepo.GetListProducts(page, limit)

	if err != nil {
		return nil, nil, err
	}

	pagination := dto.Pagination{
		CurrentPage: page,
		Limit: limit,
		TotalPages: utils.TotalPage(limit, int64(total)),
		TotalRecords: total,
	}
	return &products, &pagination, nil
}

// servide get list product by id
func (pr *productService) GetListProductById(product_id int) (model.Product, error) {
	return pr.Repo.ProductRepo.GetListProductById(product_id)
}

// service add product
func (pr *productService) AddProduct(product *model.Product) error {
	_, err := pr.Repo.CategoryRepo.GetListCategoryById(product.CategoryId)
	if err != nil {
		return errors.New("category_id is invalid or does not exist")
	}

	_, err = pr.Repo.UserRepo.GetListUserById(product.UpdatedBy)
	if err != nil {
		return errors.New("update_id (user_id) is invalid or does not exist")
	}

	err = pr.Repo.ProductRepo.AddProduct(product)
	if err != nil {
		return err
	}
	
	return nil
}

// service update product by ID
func (pr *productService) UpdateProduct(product_id int, product *model.Product) error {
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

	_, err := pr.Repo.CategoryRepo.GetListCategoryById(product.CategoryId)
	if err != nil {
		return errors.New("category_id is invalid or does not exist")
	}

	_, err = pr.Repo.UserRepo.GetListUserById(product.UpdatedBy)
	if err != nil {
		return errors.New("update_id (user_id) is invalid or does not exist")
	}

	err = pr.Repo.ProductRepo.UpdateProduct(product_id, product)
	if err != nil {
		return err
	}

	return nil
}

// service delete product by ID
func (pr *productService) DeleteProduct(product_id int) error {
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
	return pr.Repo.ProductRepo.DeleteProduct(product_id)
}