package domain

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Fullname string `json:"fullname" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	RegisterResponse struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}
)
