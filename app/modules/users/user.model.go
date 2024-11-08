package users

import (
	"time"

	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        `json:"-"`
	ID                uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName         string         `gorm:"not null" json:"first_name"`
	LastName          string         `gorm:"not null" json:"last_name"`
	MiddleName        string         `json:"middle_name"`
	Email             string         `gorm:"uniqueIndex;not null" json:"email"`
	EmailVerifiedAt   *time.Time     `json:"email_verified_at"`
	Password          string         `gorm:"not null" json:"password"`
	Username          string         `gorm:"uniqueIndex;not null" json:"username"`
	Country           string         `json:"country"`
	City              string         `json:"city"`
	ProfilePictureUrl *string        `gorm:"unique" json:"profile_picture_url"`
	CreatedAt         time.Time      `gorm:"not null;type:TIMESTAMP;" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"not null;type:TIMESTAMP;" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"type:TIMESTAMP; index"`
}

func (u *User) FullName() string {
	if u.MiddleName == "" {
		return u.FirstName + " " + u.LastName
	}
	return u.FirstName + " " + u.MiddleName + " " + u.LastName
}

func (u *User) IsEmailVerified() bool {
	return u.EmailVerifiedAt != nil
}

func (u *User) IsPasswordCorrect(password string) bool {
	return utils.VerifyDataHash(password, u.Password)
}
