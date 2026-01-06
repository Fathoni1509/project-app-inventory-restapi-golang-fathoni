package dto

type SaleCreateRequest struct {
	// SaleId        int     `json:"sale_id" validate:"required,gt=0"`
	UserId    int     `json:"user_id" validate:"required,gt=0"`
	ProductId int     `json:"product_id" validate:"required,gt=0"`
	Items     int     `json:"items" validate:"required,gt=0"`
	// Price     float32 `json:"price" validate:"required,gt=0"`
	// Total     float32 `json:"total" validate:"required,gt=0"`
}

type SaleUpdateRequest struct {
	// SaleId        int     `json:"sale_id" validate:"required,gt=0"`
	UserId    *int     `json:"user_id"`
	// ProductId *int     `json:"product_id"`
	Items     *int     `json:"items"`
	// Price     *float32 `json:"price"`
	// Total     *float32 `json:"total"`
}

type SaleResponse struct {
	SaleId    int     `json:"sale_id"`
	UserId    int     `json:"user_id"`
	ProductId int     `json:"product_id"`
	Items     int     `json:"items"`
	Price     float32 `json:"price"`
	Total     float32 `json:"total"`
}
