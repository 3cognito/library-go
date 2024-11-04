package otp

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type otpRepo struct {
	db *gorm.DB
}

type OtpRepoInterface interface {
	CreateOtp(otp *Otp) error
	GetOtpByUseCase(userId uuid.UUID, useCase string) (*Otp, error)
	SaveOtp(otp *Otp) error
}

type otpService struct {
	repo OtpRepoInterface
}

type OtpServiceInterface interface {
	CreateOtp(userId uuid.UUID, useCase UseCase, expiresAt time.Time) (*Otp, error)
	GetOtpByUseCase(userId uuid.UUID, useCase UseCase) (*Otp, error)
	InValidateOtp(otp *Otp) error
}

type UseCase string

const (
	EmailVerifcation UseCase = "email_verification"
)
