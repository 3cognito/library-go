package books

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewBookRepo(db *gorm.DB) BookRepoInterface {
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) CreateBook(book *Book) error {
	return b.db.Create(book).Error
}

func (b *bookRepo) GetBookByID(id uint) (*Book, error) {
	var book Book
	err := b.db.First(&book, id).Error
	return &book, err
}

func (b *bookRepo) Save(book *Book) error {
	return b.db.Save(book).Error
}

func (b *bookRepo) GetAuthorBooks(authorID uuid.UUID) ([]Book, error) {
	var books []Book
	err := b.db.Where("author_id = ?", authorID).Find(&books).Error
	return books, err
}

func (b *bookRepo) GetBooksByPublisher(publisher string) ([]Book, error) {
	var books []Book
	err := b.db.Where("publisher = ?", publisher).Find(&books).Error
	return books, err
}
