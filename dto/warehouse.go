package dto

type WarehouseCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Location    string `json:"location" validate:"required"`
}

type WarehouseUpdateRequest struct {
	Name        *string `json:"name"`
	Location    *string `json:"location"`
}