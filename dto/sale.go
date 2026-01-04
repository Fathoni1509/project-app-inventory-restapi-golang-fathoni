package dto

type SaleRequest struct {
	SaleId        int     `json:"sale_id" validate:"required,gt=0"`
	UserId        int     `json:"user_id" validate:"required,gt=0"`
	TotalPurchase float32 `json:"total_purchase" validate:"required,gt=0"`
}

type SaleResponse struct {
	SaleId        int     `json:"sale_id"`
	UserId        int     `json:"user_id"`
	TotalPurchase float32 `json:"total_purchase"`
}
