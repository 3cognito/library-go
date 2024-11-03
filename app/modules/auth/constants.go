package auth

import "errors"

var (
	ErrInvalidEmail         = errors.New("invalid email address")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)
