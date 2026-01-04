package dto

type InventoryRequest struct {
	InventoryId int `json:"inventory_id" validate:"required,gt=0"`
	ProductId   int `json:"product_id" validate:"required,gt=0"`
	ShelveId    int `json:"shelve_id" validate:"required,gt=0"`
	Quantity    int `json:"quantity" validate:"required,gt=0"`
}

type InventoryResponse struct {
	InventoryId int `json:"inventory_id"`
	ProductId   int `json:"product_id"`
	ShelveId    int `json:"shelve_id"`
	Quantity    int `json:"quantity"`
}
