package dto

type CategoryCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
}

type CategoryUpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
