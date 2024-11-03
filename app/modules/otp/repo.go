package otp

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewOtpRepo(db *gorm.DB) OtpRepoInterface {
	return &otpRepo{
		db: db,
	}
}

func (o *otpRepo) CreateOtp(otp *Otp) error {
	return o.db.Create(otp).Error
}

func (o *otpRepo) GetOtpByUseCase(userId uuid.UUID, useCase string) (*Otp, error) {
	var otp Otp
	err := o.db.Where("user_id = ? AND use_case = ? AND expires_at > ?", userId, useCase, time.Now()).First(&otp).Error
	return &otp, err
}
