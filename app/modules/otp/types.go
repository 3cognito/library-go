package otp

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type otpRepo struct {
	db *gorm.DB
}

type OtpRepoInterface interface {
	CreateOtp(otp *Otp) error
	GetOtpByUseCase(userId uuid.UUID, useCase string) (*Otp, error)
}

type UseCase string

const (
	EmailVerifcation UseCase = "email_verification"
)
