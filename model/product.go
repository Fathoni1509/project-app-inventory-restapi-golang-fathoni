package model

type Product struct {
	ProductId     int     `json:"product_id"`
	Name          string  `json:"name"`
	CategoryId    int     `json:"category_id"`
	Category      string  `json:"category"`
	PurchasePrice float32 `json:"purchase_price"`
	SellPrice     float32 `json:"sell_price"`
	UpdatedBy     int     `json:"updated_by"`
	Username      string  `json:"username"`
	Model
}
