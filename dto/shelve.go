package dto

type ShelveCreateRequest struct {
	// ShelveId    int    `json:"shelve_id" validate:"required,gt=0"`
	WarehouseId int    `json:"warehouse_id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=3"`
}

type ShelveUpdateRequest struct {
	// ShelveId    int    `json:"shelve_id"`
	WarehouseId *int    `json:"warehouse_id"`
	Name        *string `json:"name"`
}

type ShelveResponse struct {
	ShelveId    int    `json:"shelve_id"`
	WarehouseId int    `json:"warehouse_id"`
	Name        string `json:"name"`
}
