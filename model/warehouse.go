package model

type Warehouse struct {
	WarehouseId int    `json:"warehouse_id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Model
}
