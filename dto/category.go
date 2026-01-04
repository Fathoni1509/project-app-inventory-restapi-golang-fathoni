package dto

type CategoryCreateRequest struct {
	// CategoryId  int    `json:"category_id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
}

type CategoryUpdateRequest struct {
	// CategoryId  int    `json:"category_id" validate:"required,gt=0"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CategoryResponse struct {
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
