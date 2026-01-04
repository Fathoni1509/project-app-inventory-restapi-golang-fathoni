package model

type Category struct {
	Model
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
