package dto

type CategoryRequest struct {
	CategoryId  int    `json:"category_id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
}

type CategoryResponse struct {
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
