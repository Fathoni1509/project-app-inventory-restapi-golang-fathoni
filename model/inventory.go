package model

import "time"

type Inventory struct {
	InventoryId int       `json:"inventory_id"`
	ProductId   int       `json:"product_id"`
	ShelveId    int       `json:"shelve_id"`
	LastUpdated time.Time `json:"last_updated"`
	Quantity    int       `json:"quantity"`
	DeletedAt *time.Time `json:"deleted_at"`
}
