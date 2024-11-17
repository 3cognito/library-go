package books

import (
	"gorm.io/gorm"
)

func NewDeletedBookRepo(db *gorm.DB) DeletedBookRepoInterface {
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) CreateEntry(book *DeletedBook) error {
	return b.db.Create(book).Error
}
