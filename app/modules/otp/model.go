package otp

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Value      string         `gorm:"not null;index:idx_otp_value_user_case" json:"value"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null;index:idx_otp_value_user_case" json:"user_id"`
	UseCase    string         `gorm:"not null;index:idx_otp_value_user_case" json:"use_case"`
	ExpiresAt  time.Time      `gorm:"not null;type:TIMESTAMP;" json:"expires_at"`
	CreatedAt  time.Time      `gorm:"not null;type:TIMESTAMP;" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"not null;type:TIMESTAMP;" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"type:TIMESTAMP; index"`
}

func (o *Otp) IsExpired() bool {
	return o.ExpiresAt.Before(time.Now())
}
