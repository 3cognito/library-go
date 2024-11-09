package middlewares

import "errors"

var (
	ErrAuthRequired             = errors.New("authorization required")
	ErrInvalidToken             = errors.New("invalid authorization token")
	ErrAccountNotFoundOrDeleted = errors.New("account not found or deleted")
	ErrEmailNotVerified         = errors.New("email not verified")
)
