package dto

type CreateUserRequest struct {
	Name  string `json:"name" example:"John"`
	Email string `json:"email" example:"john@example.com"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
