package books

import (
	"gorm.io/gorm"
)

func NewDeletedBookRepo(db *gorm.DB) DeletedBookRepoInterface {
	return &deletedBookRepo{
		db: db,
	}
}

func (b *deletedBookRepo) CreateEntry(book *DeletedBook) error {
	return b.db.Create(book).Error
}
