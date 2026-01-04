package model

type Shelve struct {
	Model
	ShelveId    int    `json:"shelve_id"`
	WarehouseId int    `json:"warehouse_id"`
	Name        string `json:"name"`
}
