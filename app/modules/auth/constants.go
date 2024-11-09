package auth

import (
	"errors"

	"github.com/3cognito/library/app/utils"
)

var (
	ErrInvalidEmail         = errors.New("invalid email address")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrOtpExpiredOrInvalid  = errors.New("otp expired or invalid")
	ErrAccountNotFound      = errors.New("account not found")
)

var (
	ADayHence = utils.TimeNow().AddDate(0, 0, 1)
)
