package repository

import (
	"project-app-inventory-restapi-golang-fathoni/database"

	"go.uber.org/zap"
)

type Repository struct {
	CategoryRepo  CategoryRepository
	WarehouseRepo WarehouseRepository
	ShelveRepo    ShelveRepository
	UserRepo      UserRepository
	ProductRepo   ProductRepository
	SaleRepo      SaleRepository
	ReportRepo    ReportRepository
	DB            database.PgxIface
}

func NewRepository(db database.PgxIface, log *zap.Logger) Repository {
	return Repository{
		CategoryRepo:  NewCategoryRepository(db),
		WarehouseRepo: NewWarehouseRepository(db),
		ShelveRepo:    NewShelveRepository(db),
		UserRepo:      NewUserRepository(db),
		ProductRepo:   NewProductRepository(db, log),
		SaleRepo:      NewSaleRepository(db, log),
		ReportRepo:    NewReportRepository(db),
		DB:            db,
	}
}
