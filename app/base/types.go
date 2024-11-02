package base

import (
	"github.com/3cognito/library/app/config"
	"gorm.io/gorm"
)

type base struct {
	configs config.Config
	db      *gorm.DB
}

type appControllers struct {
}
