package utils

import "errors"

const (
	INVALID_INPUT = "invalid input"
	GENERIC_ERROR = "bad request"
)

var (
	ErrEmailAlreadyExists             = errors.New("email already exists")
	ErrUsernameAlreadyExists          = errors.New("username already exists")
	ErrProfilePictureUrlAlreadyExists = errors.New("profile picture url already exists")
)
