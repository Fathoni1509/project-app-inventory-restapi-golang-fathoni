package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/model"
)

type WarehouseRepository interface {
	GetListWarehouses() ([]model.Warehouse, error)
	GetListWarehouseById(warehouse_id int) (model.Warehouse, error)
	AddWarehouse(warehouse *model.Warehouse) error
	UpdateWarehouse(warehouse_id int, warehouse *model.Warehouse) error
	DeleteWarehouse(warehouse_id int) error
}

type warehouseRepository struct {
	db database.PgxIface
}

func NewWarehouseRepository(db database.PgxIface) WarehouseRepository {
	return &warehouseRepository{db: db}
}

// get list warehouses
func (repo *warehouseRepository) GetListWarehouses() ([]model.Warehouse, error) {
	query := `SELECT warehouse_id, name, location, created_at, updated_at, deleted_at FROM warehouses WHERE deleted_at IS NULL ORDER BY warehouse_id`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listWarehouses []model.Warehouse
	var list model.Warehouse
	for rows.Next() {
		err := rows.Scan(&list.WarehouseId, &list.Name, &list.Location, &list.CreatedAt, &list.UpdatedAt, &list.DeletedAt)
		if err != nil {
			return nil, err
		}
		listWarehouses = append(listWarehouses, list)
	}

	return listWarehouses, nil
}

// get list warehouse by ID
func (repo *warehouseRepository) GetListWarehouseById(warehouse_id int) (model.Warehouse, error) {
	query := `SELECT warehouse_id, name, location, created_at, updated_at, deleted_at 
		FROM warehouses
		WHERE deleted_at IS NULL AND warehouse_id=$1`

	var warehouse model.Warehouse

	err := repo.db.QueryRow(context.Background(), query, warehouse_id).Scan(&warehouse.WarehouseId, &warehouse.Name, &warehouse.Location, &warehouse.CreatedAt, &warehouse.UpdatedAt, &warehouse.DeletedAt)

	if err != nil {
		return model.Warehouse{}, err
	}

	return warehouse, nil
}

// add warehouse
func (repo *warehouseRepository) AddWarehouse(warehouse *model.Warehouse) error {
	query := `INSERT INTO warehouses (name, location, created_at, updated_at) VALUES
	($1, $2, NOW(), NOW()) RETURNING warehouse_id`

	_, err := repo.db.Exec(context.Background(), query,
		warehouse.Name,
		warehouse.Location,
	)

	if err != nil {
		return err
	}

	return nil
}

// update category by ID
func (repo *warehouseRepository) UpdateWarehouse(warehouse_id int, warehouse *model.Warehouse) error {
	query := `UPDATE warehouses
		SET name=$1, location=$2, updated_at=NOW()
		WHERE deleted_at IS NULL AND warehouse_id=$3`

	commandTag, err := repo.db.Exec(context.Background(), query,
		warehouse.Name,
		warehouse.Location,
		warehouse_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("warehouse not found")
	}

	return nil
}

// delete category by ID
func (repo *warehouseRepository) DeleteWarehouse(warehouse_id int) error {
	query := `UPDATE warehouses
		SET deleted_at=NOW()
		WHERE warehouse_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		warehouse_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("warehouse not found")
    }

	return nil
}
