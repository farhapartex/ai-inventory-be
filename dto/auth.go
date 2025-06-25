package dto

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignUpRequestDTO struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=50"`
	LastName  string `json:"last_name" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Gender    string `json:"gender" binding:"required,oneof=Male Female Other 'Prefer not to say'"`
}

type SignUpResponseDTO struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}
