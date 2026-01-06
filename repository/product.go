package repository

import (
	"context"
	"errors"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/model"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ProductRepository interface {
	GetListProducts(page, limit int) ([]model.Product, int, error)
	GetListProductById(product_id int) (model.Product, error)
	AddProduct(product *model.Product) error
	UpdateProduct(product_id int, product *model.Product) error
	DecreaseStock(tx pgx.Tx, product_id, quantity int) error
	DeleteProduct(product_id int) error
}

type productRepository struct {
	db database.PgxIface
	Logger *zap.Logger
}

func NewProductRepository(db database.PgxIface, log *zap.Logger) ProductRepository {
	return &productRepository{db: db, Logger: log}
}

// get list products
func (repo *productRepository) GetListProducts(page, limit int) ([]model.Product, int, error) {
	offset := (page - 1) * limit

	// get total data for pagination
	var total int
	countQuery := `SELECT COUNT(*) FROM products WHERE deleted_at IS NULL`
	err := repo.db.QueryRow(context.Background(), countQuery).Scan(&total)
	if err != nil {
		repo.Logger.Error("error query findall repo ", zap.Error(err))
		return nil, 0, err
	}

	query := `SELECT 
		product_id, 
		p.name AS name,
		p.category_id AS category_id,
		c.name AS category,
		p.shelve_id AS shelve_id,
		s.name AS shelve,
		purchase_price,
		sell_price,
		quantity,
		p.updated_by,
		u.username,
		p.created_at AS created_at, 
		p.updated_at AS updated_at, 
		p.deleted_at AS deleted_at
	FROM products p
	JOIN users u ON p.updated_by = u.user_id
	JOIN categories c ON p.category_id = c.category_id
	JOIN shelves s ON p.shelve_id = s.shelve_id
	WHERE p.deleted_at IS NULL 
	ORDER BY product_id
	LIMIT $1 OFFSET $2`

	rows, err := repo.db.Query(context.Background(), query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var listProducts []model.Product
	var list model.Product
	for rows.Next() {
		err := rows.Scan(&list.ProductId, &list.Name, &list.CategoryId, &list.Category, &list.ShelveId, &list.Shelve, &list.PurchasePrice, &list.SellPrice, &list.Quantity, &list.UpdatedBy, &list.Username, &list.CreatedAt, &list.UpdatedAt, &list.DeletedAt)
		if err != nil {
			return nil, 0, err
		}
		listProducts = append(listProducts, list)
	}

	return listProducts, total, nil
}

// get list product by ID
func (repo *productRepository) GetListProductById(product_id int) (model.Product, error) {
	query := `SELECT 
		product_id, 
		p.name AS name,
		p.category_id AS category_id,
		c.name AS category,
		p.shelve_id AS shelve_id,
		s.name AS shelve,
		purchase_price,
		sell_price,
		quantity,
		p.updated_by,
		u.username AS updated_by,
		p.created_at AS created_at, 
		p.updated_at AS updated_at, 
		p.deleted_at AS deleted_at
	FROM products p
	JOIN users u ON p.updated_by = u.user_id
	JOIN categories c ON p.category_id = c.category_id
	JOIN shelves s ON p.shelve_id = s.shelve_id
	WHERE p.deleted_at IS NULL AND p.product_id=$1
	ORDER BY product_id`

	var product model.Product

	err := repo.db.QueryRow(context.Background(), query, product_id).Scan(&product.ProductId, &product.Name, &product.CategoryId, &product.Category, &product.ShelveId, &product.Shelve, &product.PurchasePrice, &product.SellPrice, &product.Quantity, &product.UpdatedBy, &product.Username, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

// add product
func (repo *productRepository) AddProduct(product *model.Product) error {
	query := `INSERT INTO products (category_id, name, shelve_id, purchase_price, sell_price, quantity, updated_by, created_at, updated_at) VALUES
	($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING product_id`

	_, err := repo.db.Exec(context.Background(), query,
		product.CategoryId,
		product.Name,
		product.ShelveId,
		product.PurchasePrice,
		product.SellPrice,
		product.Quantity,
		product.UpdatedBy,
	)

	if err != nil {
		return err
	}

	return nil
}

// update product by ID
func (repo *productRepository) UpdateProduct(product_id int, product *model.Product) error {
	query := `UPDATE products
		SET category_id=$1, name=$2, shelve_id=$3, purchase_price=$4, sell_price=$5, quantity=$6, updated_by=$7, updated_at=NOW()
		WHERE deleted_at IS NULL AND product_id=$8`

	commandTag, err := repo.db.Exec(context.Background(), query,
		product.CategoryId,
		product.Name,
		product.ShelveId,
		product.PurchasePrice,
		product.SellPrice,
		product.Quantity,
		product.UpdatedBy,
		product_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("product not found")
	}

	return nil
}

// decrease stock
func (repo *productRepository) DecreaseStock(tx pgx.Tx, product_id, quantity int) error {
	query := `UPDATE products
		SET quantity = quantity - $1, updated_at=NOW()
		WHERE product_id=$2 AND deleted_at IS NULL AND quantity>=$1`
	
	commandTag, err := tx.Exec(context.Background(), query, quantity, product_id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("stock not enough")
    }

	return nil
}

// delete product by ID
func (repo *productRepository) DeleteProduct(product_id int) error {
	query := `UPDATE products
		SET deleted_at=NOW()
		WHERE product_id=$1 AND deleted_at IS NULL`

	commandTag, err := repo.db.Exec(context.Background(), query,
		product_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("product not found")
    }

	return nil
}
