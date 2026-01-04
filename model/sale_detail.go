package model

type SaleDetail struct {
	SaleDetailId  int     `json:"sale_detail_id"`
	SaleId        int     `json:"sale_id"`
	ProductId     int     `json:"product_id"`
	TotalItems    int     `json:"total_items"`
	PricePerItems float32 `json:"price_per_items"`
	SubTotal      float32 `json:"sub_total"`
}
