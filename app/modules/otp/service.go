package otp

import (
	"time"

	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
)

func NewOtpService(repo OtpRepoInterface) OtpServiceInterface {
	return &otpService{
		repo: repo,
	}
}

func (o *otpService) CreateOtp(userId uuid.UUID, useCase UseCase, expiresAt time.Time) (string, error) {
	retries := 3
	otp := &Otp{
		UserID:    userId,
		UseCase:   string(useCase),
		ExpiresAt: expiresAt,
		Value:     utils.GenerateOtp(),
	}
	err := o.repo.CreateOtp(otp)
	if err != nil {
		//TODO: might want to adjust expiry time by calculating time elapsed and adding to expiry time
		for i := 0; i < retries; i++ {
			otp.Value = utils.GenerateOtp()
			err = o.repo.CreateOtp(otp)
			if err == nil {
				break
			}
		}
	}
	return otp.Value, err
}

func (o *otpService) GetOtpByUseCase(userId uuid.UUID, useCase UseCase) (string, error) {
	otp, err := o.repo.GetOtpByUseCase(userId, string(useCase))
	if err != nil {
		return "", err
	}
	return otp.Value, nil
}
