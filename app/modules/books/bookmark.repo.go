package books

import "gorm.io/gorm"

func NewBookMarkRepo(
	db *gorm.DB,
) BookMarkRepoInterface {
	return &bookMarkRepo{
		db: db,
	}
}
