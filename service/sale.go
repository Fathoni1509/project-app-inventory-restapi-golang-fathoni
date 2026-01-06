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
		CurrentPage: page,
		Limit: limit,
		TotalPages: utils.TotalPage(limit, int64(total)),
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

// service update sale by ID

/*
func (sl *saleService) UpdateSale(sale_id int, sale *model.Sale) error {
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

	_, err := sl.Repo.UserRepo.GetListUserById(sale.UserId)
	if err != nil {
		return errors.New("user_id is invalid or does not exist")
	}

	_, err = sl.Repo.ProductRepo.GetListProductById(sale.ProductId)
	if err != nil {
		return errors.New("product_id is invalid or does not exist")
	}

	err = sl.Repo.SaleRepo.UpdateSale(sale_id, sale)
	if err != nil {
		return err
	}

	return nil
} */

func (sl *saleService) UpdateSale(saleID int, reqUpdate *model.Sale) error {
    // --- MULAI TRANSAKSI DARI AWAL ---
    tx, err := sl.Repo.DB.Begin(context.Background())
    if err != nil {
        return err
    }
    defer tx.Rollback(context.Background())

    // A. Ambil Data Sale Lama (Existing)
    // Note: Sebaiknya buat fungsi GetById yang menerima 'tx' di repository 
    // agar thread-safe, tapi pakai yang biasa dulu tidak apa-apa asal hati-hati.
    oldSale, err := sl.Repo.SaleRepo.GetListSaleById(saleID)
    if err != nil {
        return errors.New("sale not found")
    }

    // Validasi: Tidak boleh ganti produk (Terlalu kompleks logikanya)
    // Jika user mau ganti produk, suruh hapus sale dan buat baru.

    // if oldSale.ProductId != reqUpdate.ProductId {
    //      return errors.New("cannot change product type, please delete and create new sale")
    // }

    // B. Hitung Selisih Quantity (Delta)
    // Baru: 5, Lama: 2 -> Diff: 3 (Kurangi stok 3 lagi)
    // Baru: 2, Lama: 5 -> Diff: -3 (Kembalikan stok 3)
	
    diffQty := reqUpdate.Items - oldSale.Items
	reqUpdate.ProductId = oldSale.ProductId

    // C. Update Stok jika ada perubahan jumlah
    if diffQty != 0 {
        err = sl.Repo.ProductRepo.DecreaseStock(tx, reqUpdate.ProductId, diffQty)
        if err != nil {
            return errors.New("insufficient stock for update")
        }
    }

    // D. Update Harga & Total Baru
    product, _ := sl.Repo.ProductRepo.GetListProductById(reqUpdate.ProductId)
    reqUpdate.Price = product.SellPrice
    reqUpdate.Total = float32(reqUpdate.Items) * reqUpdate.Price
    
    // Pastikan field UserId terisi (pakai yang lama jika kosong di request)
    // (Logic merging biasanya sudah di handler, tapi untuk aman set disini juga ok)
    // reqUpdate.UserId = oldSale.UserId 

    // E. Update Sale Record
    err = sl.Repo.SaleRepo.UpdateSale(tx, saleID, reqUpdate)
    if err != nil {
        return err
    }

    // F. Commit
    return tx.Commit(context.Background())
}

// service delete sale by ID
func (sl *saleService) DeleteSale(sale_id int) error {
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
	return sl.Repo.SaleRepo.DeleteSale(sale_id)
}