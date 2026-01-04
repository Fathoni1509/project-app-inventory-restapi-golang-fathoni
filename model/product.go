package model

type Product struct {
	Model
	ProductId     int     `json:"product_id"`
	Name          string  `json:"name"`
	CategoryId    int     `json:"category_id"`
	PurchasePrice float32 `json:"purchase_price"`
	SellPrice     float32 `json:"sell_price"`
	UpdatedBy     string  `json:"updated_by"`
}
