package domain

import "time"

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Image     string    `json:"image"`
	Gender    string    `json:"gender"`
	Dob       time.Time `json:"dob"`
	IsPremium int       `json:"is_premium"`
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

	ProfileResponse struct {
		Id        string    `json:"id"`
		Username  string    `json:"username"`
		Fullname  string    `json:"fullname"`
		Email     string    `json:"email"`
		Image     string    `json:"image"`
		Gender    string    `json:"gender"`
		Dob       time.Time `json:"dob"`
		IsPremium int       `json:"is_premium"`
	}
)

const (
	Male        = "male"
	Female      = "female"
	Undisclosed = "undisclosed"
)
