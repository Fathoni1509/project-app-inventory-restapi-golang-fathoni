package dto

type SaleCreateRequest struct {
	UserId    int     `json:"user_id" validate:"required,gt=0"`
	ProductId int     `json:"product_id" validate:"required,gt=0"`
	Items     int     `json:"items" validate:"required,gt=0"`
}

type SaleUpdateRequest struct {
	UserId    *int     `json:"user_id"`
	Items     *int     `json:"items"`
}
