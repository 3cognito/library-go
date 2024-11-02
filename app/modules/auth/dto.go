package auth

type SignUpRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name,omitempty"`
	Username   string `json:"username" binding:"required"`
	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
}

type UserResponse struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
	Username   string `json:"username"`
	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
}

type SignUpResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
