package dto

import "time"

type InventoryCreateRequest struct {
	// InventoryId int `json:"inventory_id" validate:"required,gt=0"`
	ProductId int `json:"product_id" validate:"required,gt=0"`
	ShelveId  int `json:"shelve_id" validate:"required,gt=0"`
	Quantity  int `json:"quantity" validate:"required,gt=0"`
}

type InventoryUpdateRequest struct {
	// InventoryId *int `json:"inventory_id"`
	ProductId *int `json:"product_id"`
	ShelveId  *int `json:"shelve_id"`
	Quantity  *int `json:"quantity"`
}

type InventoryResponse struct {
	InventoryId int        `json:"inventory_id"`
	ProductId   int        `json:"product_id"`
	Product     string     `json:"product"`
	ShelveId    int        `json:"shelve_id"`
	Shelve      string     `json:"shelve"`
	LastUpdated time.Time  `json:"last_updated"`
	Quantity    int        `json:"quantity"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
