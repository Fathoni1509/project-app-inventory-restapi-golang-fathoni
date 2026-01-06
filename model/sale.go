package model

type Sale struct {
	SaleId    int     `json:"sale_id"`
	UserId    int     `json:"user_id"`
	User      string  `json:"user"`
	ProductId int     `json:"product_id"`
	Product   string     `json:"product"`
	Items     int     `json:"items"`
	Price     float32 `json:"price"`
	Total     float32 `json:"total"`
	Model
}
