package middlewares

import "errors"

var (
	ErrAuthRequired             = errors.New("authorization required")
	ErrInvalidOrExpiredToken    = errors.New("invalid or expired token")
	ErrAccountNotFoundOrDeleted = errors.New("account not found or deleted")
	ErrEmailNotVerified         = errors.New("email not verified")
)
