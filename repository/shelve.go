package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/model"
)

type ShelveRepository interface {
	GetListShelves() ([]model.Shelve, error)
	GetListShelveById(shelve_id int) (model.Shelve, error)
	AddShelve(shelve *model.Shelve) error
	UpdateShelve(shelve_id int, shelve *model.Shelve) error
	DeleteShelve(shelve_id int) error
}

type shelveRepository struct {
	db database.PgxIface
}

func NewShelveRepository(db database.PgxIface) ShelveRepository {
	return &shelveRepository{db: db}
}

// get list shelves
func (repo *shelveRepository) GetListShelves() ([]model.Shelve, error) {
	query := `SELECT 
		shelve_id, 
		s.warehouse_id, 
		s.name AS shelve_name,
		w.name AS warehouse,
		s.created_at AS created_at, 
		s.updated_at AS updated_at, 
		s.deleted_at AS deleted_at
	FROM shelves s
	JOIN warehouses w ON s.warehouse_id = w.warehouse_id
	WHERE s.deleted_at IS NULL 
	ORDER BY shelve_id`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listShelves []model.Shelve
	var list model.Shelve
	for rows.Next() {
		err := rows.Scan(&list.ShelveId, &list.WarehouseId, &list.Name, &list.Warehouse, &list.CreatedAt, &list.UpdatedAt, &list.DeletedAt)
		if err != nil {
			return nil, err
		}
		listShelves = append(listShelves, list)
	}

	return listShelves, nil
}

// get list shelve by ID
func (repo *shelveRepository) GetListShelveById(shelve_id int) (model.Shelve, error) {
	query := `SELECT 
		shelve_id,
		s.warehouse_id, 
		s.name AS shelve_name,
		w.name AS warehouse,
		s.created_at AS created_at, 
		s.updated_at AS updated_at, 
		s.deleted_at AS deleted_at
	FROM shelves s
	JOIN warehouses w ON s.warehouse_id = w.warehouse_id
	WHERE s.deleted_at IS NULL AND shelve_id=$1
	ORDER BY shelve_id`

	var shelve model.Shelve

	err := repo.db.QueryRow(context.Background(), query, shelve_id).Scan(&shelve.ShelveId, &shelve.WarehouseId, &shelve.Name, &shelve.Warehouse, &shelve.CreatedAt, &shelve.UpdatedAt, &shelve.DeletedAt)

	if err != nil {
		return model.Shelve{}, err
	}

	return shelve, nil
}

// add shelve
func (repo *shelveRepository) AddShelve(shelve *model.Shelve) error {
	query := `INSERT INTO shelves (warehouse_id, name, created_at, updated_at) VALUES
	($1, $2, NOW(), NOW()) RETURNING shelve_id`

	_, err := repo.db.Exec(context.Background(), query,
		shelve.WarehouseId,
		shelve.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

// update shelve by ID
func (repo *shelveRepository) UpdateShelve(shelve_id int, shelve *model.Shelve) error {
	query := `UPDATE shelves
		SET warehouse_id =$1, name=$2, updated_at=NOW()
		WHERE deleted_at IS NULL AND shelve_id=$3`

	commandTag, err := repo.db.Exec(context.Background(), query,
		shelve.WarehouseId,
		shelve.Name,
		shelve_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("shelve not found")
	}

	return nil
}

// delete shelve by ID
func (repo *shelveRepository) DeleteShelve(shelve_id int) error {
	query := `UPDATE shelves
		SET deleted_at=NOW()
		WHERE shelve_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		shelve_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("shelve not found")
    }

	return nil
}
