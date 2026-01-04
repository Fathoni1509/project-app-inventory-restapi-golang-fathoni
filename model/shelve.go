package model

type Shelve struct {
	ShelveId    int    `json:"shelve_id"`
	WarehouseId int    `json:"warehouse_id"`
	Warehouse   string `json:"warehouse"`
	Name        string `json:"name"`
	Model
}
