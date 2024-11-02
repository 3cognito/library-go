package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) BeginTrx() *gorm.DB {
	return u.db.Begin()
}

func (u *userRepo) CreateUser(user *User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *userRepo) GetUserByEmail(email string) (*User, error) {
	var user User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *userRepo) GetUserByUsername(username string) (*User, error) {
	var user User
	err := u.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
