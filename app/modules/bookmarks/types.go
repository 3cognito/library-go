package bookmarks

import "gorm.io/gorm"

type bookMarkRepo struct {
	db *gorm.DB
}

type BookMarkRepoInterface interface{}
