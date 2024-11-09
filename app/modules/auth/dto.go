package auth

import "github.com/google/uuid"

type SignUpRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name" `
	MiddleName string `json:"middle_name,omitempty"`
	Username   string `json:"username"`
	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
}

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name,omitempty"`
	Username   string    `json:"username"`
	Country    string    `json:"country,omitempty"`
	City       string    `json:"city,omitempty"`
}

type LoggedInResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyEmailRequest struct {
	Otp    string    `json:"otp"`
	UserID uuid.UUID `json:"user_id"`
}
