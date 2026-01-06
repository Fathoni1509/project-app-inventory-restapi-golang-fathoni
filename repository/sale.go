package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/model"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type SaleRepository interface {
	GetListSales(page, limit int) ([]model.Sale, int, error)
	GetListSaleById(sale_id int) (model.Sale, error)
	AddSale(tx pgx.Tx, sale *model.Sale) error
	UpdateSale(tx pgx.Tx, sale_id int, sale *model.Sale) error
	DeleteSale(sale_id int) error
}

type saleRepository struct {
	db database.PgxIface
	Logger *zap.Logger
}

func NewSaleRepository(db database.PgxIface, log *zap.Logger) SaleRepository {
	return &saleRepository{db: db, Logger: log}
}

// get list sales
func (repo *saleRepository) GetListSales(page, limit int) ([]model.Sale, int, error) {
	offset := (page - 1) * limit

	// get total data for pagination
	var total int
	countQuery := `SELECT COUNT(*) FROM sales WHERE deleted_at IS NULL`
	err := repo.db.QueryRow(context.Background(), countQuery).Scan(&total)
	if err != nil {
		repo.Logger.Error("error query findall repo ", zap.Error(err))
		return nil, 0, err
	}

	query := `SELECT 
		sale_id,
		s.user_id,
		u.username AS user,
		s.product_id,
		p.name,
		items,
		price,
		total,
		s.created_at AS created_at, 
		s.updated_at AS updated_at, 
		s.deleted_at AS deleted_at
	FROM sales s
	JOIN users u ON s.user_id = u.user_id
	JOIN products p ON s.product_id = p.product_id
	WHERE s.deleted_at IS NULL 
	ORDER BY sale_id
	LIMIT $1 OFFSET $2`

	rows, err := repo.db.Query(context.Background(), query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var listSales []model.Sale
	var list model.Sale
	for rows.Next() {
		err := rows.Scan(&list.SaleId, &list.UserId, &list.User, &list.ProductId, &list.Product, &list.Items, &list.Price, &list.Total, &list.CreatedAt, &list.UpdatedAt, &list.DeletedAt)
		if err != nil {
			return nil, 0, err
		}
		listSales = append(listSales, list)
	}

	return listSales, total, nil
}

// get list sale by ID
func (repo *saleRepository) GetListSaleById(sale_id int) (model.Sale, error) {
	query := `SELECT 
		sale_id,
		s.user_id,
		u.username AS user,
		s.product_id,
		p.name,
		items,
		p.sell_price,
		total,
		s.created_at AS created_at, 
		s.updated_at AS updated_at, 
		s.deleted_at AS deleted_at
	FROM sales s
	JOIN users u ON s.user_id = u.user_id
	JOIN products p ON s.product_id = p.product_id
	WHERE s.deleted_at IS NULL AND sale_id=$1
	ORDER BY sale_id`

	var sale model.Sale

	err := repo.db.QueryRow(context.Background(), query, sale_id).Scan(&sale.SaleId, &sale.UserId, &sale.User, &sale.ProductId, &sale.Product, &sale.Items, &sale.Price, &sale.Total, &sale.CreatedAt, &sale.UpdatedAt, &sale.DeletedAt)

	if err != nil {
		return model.Sale{}, err
	}

	return sale, nil
}

// add sale
func (repo *saleRepository) AddSale(tx pgx.Tx, sale *model.Sale) error {
	query := `INSERT INTO sales (user_id, product_id, items, price, total, created_at, updated_at) VALUES
	($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING sale_id`

	_, err := tx.Exec(context.Background(), query,
		sale.UserId,
		sale.ProductId,
		sale.Items,
		sale.Price,
		sale.Total,
	)

	if err != nil {
		return err
	}

	return nil
}

// update sale by ID
func (repo *saleRepository) UpdateSale(tx pgx.Tx, sale_id int, sale *model.Sale) error {
	query := `UPDATE sales
		SET user_id=$1, product_id=$2, items=$3, price=$4, total=$5, updated_at=NOW()
		WHERE deleted_at IS NULL AND sale_id=$6`

	commandTag, err := tx.Exec(context.Background(), query,
		sale.UserId,
		sale.ProductId,
		sale.Items,
		sale.Price,
		sale.Total,
		sale_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("sale not found")
	}

	return nil
}

// delete sale by ID
func (repo *saleRepository) DeleteSale(sale_id int) error {
	query := `UPDATE sales
		SET deleted_at=NOW()
		WHERE sale_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		sale_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("sale not found")
    }

	return nil
}
