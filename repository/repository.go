package repository

import (
	"project-app-inventory-restapi-golang-fathoni/database"

	"go.uber.org/zap"
)

type Repository struct {
	CategoryRepo CategoryRepository
	WarehouseRepo WarehouseRepository
}

func NewRepository(db database.PgxIface, log *zap.Logger) Repository {
	return Repository{
		CategoryRepo: NewCategoryRepository(db),
		WarehouseRepo: NewWarehouseRepository(db),
	}
}