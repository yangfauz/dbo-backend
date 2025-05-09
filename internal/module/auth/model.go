package auth

// Request
type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

// Response
type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

type RegisterResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}
