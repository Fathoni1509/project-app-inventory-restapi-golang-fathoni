package service

import (
	// "errors"
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
	"project-app-inventory-restapi-golang-fathoni/repository"
	"project-app-inventory-restapi-golang-fathoni/utils"
)

type SaleService interface {
	GetListSales(page, limit int) (*[]model.Sale, *dto.Pagination, error)
	GetListSaleById(sale_id int) (model.Sale, error)
	AddSale(sale *model.Sale) error
	// UpdateSale(sale_id int, sale *model.Sale) error
	UpdateSale(saleID int, reqUpdate *model.Sale) error
	DeleteSale(sale_id int) error
}

type saleService struct {
	Repo repository.Repository
}

func NewSaleService(repo repository.Repository) SaleService {
	return &saleService{Repo: repo}
}

// service get list sales
func (sl *saleService) GetListSales(page, limit int) (*[]model.Sale, *dto.Pagination, error) {
	sales, total, err := sl.Repo.SaleRepo.GetListSales(page, limit)

	if err != nil {
		return nil, nil, err
	}

	pagination := dto.Pagination{
		CurrentPage:  page,
		Limit:        limit,
		TotalPages:   utils.TotalPage(limit, int64(total)),
		TotalRecords: total,
	}
	return &sales, &pagination, nil
}

// servide get list sale by id
func (sl *saleService) GetListSaleById(sale_id int) (model.Sale, error) {
	return sl.Repo.SaleRepo.GetListSaleById(sale_id)
}

// service add sale
func (sl *saleService) AddSale(sale *model.Sale) error {
	_, err := sl.Repo.UserRepo.GetListUserById(sale.UserId)
	if err != nil {
		return errors.New("user_id is invalid or does not exist")
	}

	product, err := sl.Repo.ProductRepo.GetListProductById(sale.ProductId)
	if err != nil {
		return errors.New("product_id is invalid or does not exist")
	}

	// Set Harga dari Database (Security)
	sale.Price = product.SellPrice
	sale.Total = sale.Price * float32(sale.Items)

	tx, err := sl.Repo.DB.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	// decrease stock
	err = sl.Repo.ProductRepo.DecreaseStock(tx, sale.ProductId, sale.Items)
	if err != nil {
		return err
	}

	err = sl.Repo.SaleRepo.AddSale(tx, sale)
	if err != nil {
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func (sl *saleService) UpdateSale(saleID int, reqUpdate *model.Sale) error {
	// start transaction
	tx, err := sl.Repo.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// get old sale
	oldSale, err := sl.Repo.SaleRepo.GetListSaleById(saleID)
	if err != nil {
		return errors.New("sale not found")
	}

	// count difference quantity

	diffQty := reqUpdate.Items - oldSale.Items
	reqUpdate.ProductId = oldSale.ProductId

	// update stock if there are changes
	if diffQty != 0 {
		err = sl.Repo.ProductRepo.DecreaseStock(tx, reqUpdate.ProductId, diffQty)
		if err != nil {
			return errors.New("insufficient stock for update")
		}
	}

	// update price and total
	product, _ := sl.Repo.ProductRepo.GetListProductById(reqUpdate.ProductId)
	reqUpdate.Price = product.SellPrice
	reqUpdate.Total = float32(reqUpdate.Items) * reqUpdate.Price

	// update sale record
	err = sl.Repo.SaleRepo.UpdateSale(tx, saleID, reqUpdate)
	if err != nil {
		return err
	}

	// commit
	return tx.Commit(context.Background())
}

// service delete sale by ID
func (sl *saleService) DeleteSale(sale_id int) error {
	return sl.Repo.SaleRepo.DeleteSale(sale_id)
}
