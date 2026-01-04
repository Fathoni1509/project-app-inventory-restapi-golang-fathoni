package model

type Sale struct {
	Model
	SaleId        int     `json:"sale_id"`
	UserId        int     `json:"user_id"`
	TotalPurchase float32 `json:"total_purchase"`
}
