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

func (o *otpService) CreateOtp(userId uuid.UUID, useCase UseCase, expiresAt time.Time) (*Otp, error) {
	//invalidate any existing otp for the user and use case
	existingOtp, existingErr := o.GetOtpByUseCase(userId, useCase)
	if existingErr == nil {
		o.InValidateOtp(existingOtp)
	}

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
	return otp, err
}

func (o *otpService) GetOtpByUseCase(userId uuid.UUID, useCase UseCase) (*Otp, error) {
	otp, err := o.repo.GetOtpByUseCase(userId, string(useCase))
	if err != nil {
		return otp, err
	}
	return otp, nil
}

func (o *otpService) InValidateOtp(otp *Otp) error {
	otp.InValidate()
	return o.repo.SaveOtp(otp)
}
