package dto

type SaleDetailRequest struct {
	SaleDetailId  int     `json:"sale_detail_id" validate:"required,gt=0"`
	SaleId        int     `json:"sale_id" validate:"required,gt=0"`
	ProductId     int     `json:"product_id" validate:"required,gt=0"`
	TotalItems    int     `json:"total_items" validate:"required,gt=0"`
	PricePerItems float32 `json:"price_per_items" validate:"required,gt=0"`
	SubTotal      float32 `json:"sub_total" validate:"required,gt=0"`
}

type SaleDetailResponse struct {
	SaleDetailId  int     `json:"sale_detail_id"`
	SaleId        int     `json:"sale_id"`
	ProductId     int     `json:"product_id"`
	TotalItems    int     `json:"total_items"`
	PricePerItems float32 `json:"price_per_items"`
	SubTotal      float32 `json:"sub_total"`
}
