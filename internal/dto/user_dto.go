package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	Username     string `json:"username" validate:"omitempty,min=3"`
	Email        string `json:"email" validate:"omitempty,email"`
	LastPassword string `json:"last_password" validate:"omitempty,min=6"`
	NewPassword  string `json:"new_password" validate:"omitempty,min=6"`
	ProfileImg   string `json:"profile_img" validate:"omitempty"`
}
