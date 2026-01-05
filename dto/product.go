package dto

type ProductCreateRequest struct {
	// ProductId     int     `json:"product_id" validate:"required,gt=0"`
	Name          string  `json:"name" validate:"required,min=3"`
	CategoryId    int     `json:"category_id" validate:"required,gt=0"`
	PurchasePrice float32 `json:"purchase_price" validate:"required,gt=0"`
	SellPrice     float32 `json:"sell_price" validate:"required,gt=0"`
	UpdatedBy     int     `json:"updated_by" validate:"required"`
}

type ProductUpdateRequest struct {
	// ProductId     *int     `json:"product_id"`
	Name          *string  `json:"name"`
	CategoryId    *int     `json:"category_id"`
	PurchasePrice *float32 `json:"purchase_price"`
	SellPrice     *float32 `json:"sell_price"`
	UpdatedBy     *int     `json:"updated_by" validate:"required"`
}

type ProductResponse struct {
	ProductId     int     `json:"product_id"`
	Name          string  `json:"name"`
	CategoryId    int     `json:"category_id"`
	PurchasePrice float32 `json:"purchase_price"`
	SellPrice     float32 `json:"sell_price"`
	UpdatedBy     int     `json:"updated_by"`
}
