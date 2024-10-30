package base

import (
	"github.com/3cognito/library/app/config"
	"gorm.io/gorm"
)

func New(configs config.Config, db *gorm.DB) *applicationBase {
	return &applicationBase{
		configs: configs,
		db:      db,
	}
}

func (b *applicationBase) LoadControllers() appControllers {
	var c appControllers

	return c
}

func (b *applicationBase) LoadAdminControllers() adminControllers {
	var c adminControllers

	return c
}
