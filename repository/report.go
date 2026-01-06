package repository

import (
	"context"
	"project-app-inventory-restapi-golang-fathoni/database"
	"project-app-inventory-restapi-golang-fathoni/dto"
)

type ReportRepository interface {
	GetListReports() ([]dto.ReportResponse, error)
	GetListMinStocks() ([]dto.MinStockResponse, error)
}

type reportRepository struct {
	db database.PgxIface
}

func NewReportRepository(db database.PgxIface) ReportRepository {
	return &reportRepository{db: db}
}

// get list reports
func (repo *reportRepository) GetListReports() ([]dto.ReportResponse, error) {
	query := `SELECT  
		p.name AS name,
		p.quantity AS remain,
		SUM(s.items) AS sold,
		s.price AS sell_price,
		SUM(s.total) AS sales,
		(SUM(s.total) - (p.purchase_price*(SUM(s.items)))) AS income
	FROM sales s
	JOIN products p ON s.product_id = p.product_id
	WHERE s.deleted_at IS NULL
	GROUP BY p.product_id, p.name, s.price
	ORDER BY income DESC`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var response []dto.ReportResponse
	var list dto.ReportResponse
	for rows.Next() {
		err := rows.Scan(&list.Name, &list.Remain, &list.Sold, &list.SellPrice, &list.Sales, &list.Income)
		if err != nil {
			return nil, err
		}
		response = append(response, list)
	}

	return response, nil
}

// get list minstock
func (repo *reportRepository) GetListMinStocks() ([]dto.MinStockResponse, error) {
	query := `SELECT  
		name AS name,
		quantity AS stock
	FROM products
	WHERE deleted_at IS NULL AND quantity < 5
	ORDER BY stock`

	rows, err := repo.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var response []dto.MinStockResponse
	var list dto.MinStockResponse
	for rows.Next() {
		err := rows.Scan(&list.Name, &list.Stock)
		if err != nil {
			return nil, err
		}
		response = append(response, list)
	}

	return response, nil
}