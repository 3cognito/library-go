package base

import (
	"github.com/3cognito/library/app/config"
	"gorm.io/gorm"
)

func New(configs config.Config, db *gorm.DB) *base {
	return &base{
		configs: configs,
		db:      db,
	}
}

func (b *base) LoadControllers() appControllers {
	var c appControllers

	c.AuthC = b.WithAuthController()

	return c
}
