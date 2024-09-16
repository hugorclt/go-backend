package dto

type RegisterDTO struct {
	Firstname string `json:"firstname" validate:"required,min=2,max=50"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}
