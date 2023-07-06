package authdto

type AuthResponseDTO struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}
