package auth

import (
	"errors"

	"github.com/3cognito/library/app/utils"
)

var (
	ErrInvalidEmail         = errors.New("invalid email address")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)

var (
	ADayHence = utils.TimeNow().AddDate(0, 0, 1)
)
