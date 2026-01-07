package dto

type ShelveCreateRequest struct {
	WarehouseId int    `json:"warehouse_id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=3"`
}

type ShelveUpdateRequest struct {
	WarehouseId *int    `json:"warehouse_id"`
	Name        *string `json:"name"`
}