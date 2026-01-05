package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/dto"
	"project-app-inventory-restapi-golang-fathoni/model"
)

type InventoryRepository interface {
	GetListInventorys() ([]dto.InventoryResponse, error)
	GetListInventoryById(inventory_id int) (dto.InventoryResponse, error)
	AddInventory(inventory *model.Inventory) error
	UpdateInventory(inventory_id int, inventory *model.Inventory) error
	DeleteInventory(inventory_id int) error
}

type inventoryRepository struct {
	db database.PgxIface
}

func NewInventoryRepository(db database.PgxIface) InventoryRepository {
	return &inventoryRepository{db: db}
}

// get list inventorys
func (repo *inventoryRepository) GetListInventorys() ([]dto.InventoryResponse, error) {
	query := `SELECT 
		inventory_id, 
		i.product_id AS product_id,
		p.name AS product,
		i.shelve_id AS shelve_id,
		s.name AS shelve,
		last_updated,
		quantity
	FROM inventory i
	JOIN products p ON i.product_id = p.product_id
	JOIN shelves s ON i.shelve_id = s.shelve_id
	WHERE i.deleted_at IS NULL
	ORDER BY inventory_id`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var response []dto.InventoryResponse
	var list dto.InventoryResponse
	for rows.Next() {
		err := rows.Scan(&list.InventoryId, &list.ProductId, &list.Product, &list.ShelveId, &list.Shelve, &list.LastUpdated, &list.Quantity)
		if err != nil {
			return nil, err
		}
		response = append(response, list)
	}

	return response, nil
}

// get list inventory by ID
func (repo *inventoryRepository) GetListInventoryById(inventory_id int) (dto.InventoryResponse, error) {
	query := `SELECT 
		inventory_id, 
		i.product_id AS product_id,
		p.name AS product,
		i.shelve_id AS shelve_id,
		s.name AS shelve,
		last_updated,
		quantity
	FROM inventory i
	JOIN products p ON i.product_id = p.product_id
	JOIN shelves s ON i.shelve_id = s.shelve_id
	WHERE i.deleted_at IS NULL AND inventory_id=$1
	ORDER BY inventory_id`

	var response dto.InventoryResponse

	err := repo.db.QueryRow(context.Background(), query, inventory_id).Scan(&response.InventoryId, &response.ProductId, &response.Product, &response.ShelveId, &response.Shelve, &response.LastUpdated, &response.Quantity)

	if err != nil {
		return dto.InventoryResponse{}, err
	}

	return response, nil
}

// add inventory
func (repo *inventoryRepository) AddInventory(inventory *model.Inventory) error {
	query := `INSERT INTO inventory (product_id, shelve_id, last_updated, quantity) VALUES
	($1, $2, NOW(), $3) RETURNING inventory_id`

	_, err := repo.db.Exec(context.Background(), query,
		inventory.ProductId,
		inventory.ShelveId,
		inventory.Quantity,
	)

	if err != nil {
		return err
	}

	return nil
}

// update inventory by ID
func (repo *inventoryRepository) UpdateInventory(inventory_id int, inventory *model.Inventory) error {
	query := `UPDATE inventory
		SET product_id=$1, shelve_id=$2, last_updated=NOW(), quantity=$3
		WHERE deleted_at IS NULL AND inventory_id=$4`

	commandTag, err := repo.db.Exec(context.Background(), query,
		inventory.ProductId,
		inventory.ShelveId,
		inventory.Quantity,
		inventory_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("inventory not found")
	}

	return nil
}

// delete inventory by ID
func (repo *inventoryRepository) DeleteInventory(inventory_id int) error {
	query := `UPDATE inventory
		SET deleted_at=NOW()
		WHERE inventory_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		inventory_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("inventory not found")
    }

	return nil
}
