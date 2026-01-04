package dto

type WarehouseCreateRequest struct {
	// WarehouseId int    `json:"warehouse_id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=3"`
	Location    string `json:"location" validate:"required"`
}

type WarehouseUpdateRequest struct {
	// WarehouseId int    `json:"warehouse_id" validate:"required,gt=0"`
	Name        *string `json:"name"`
	Location    *string `json:"location"`
}

type WarehouseResponse struct {
	WarehouseId int    `json:"warehouse_id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
}
