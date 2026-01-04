package model

type Warehouse struct {
	Model
	WarehouseId int    `json:"warehouse_id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
}
