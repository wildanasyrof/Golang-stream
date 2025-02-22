package dto

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
