package books

import (
	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewBookRepo(db *gorm.DB) BookRepoInterface {
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) CreateBook(book *Book) error {
	res := b.db.Create(book)
	return utils.CheckUniqueConstrainstErr(res.Error)
}

func (b *bookRepo) GetBookByID(id uuid.UUID) (*Book, error) {
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

func (b *bookRepo) GetAuthorBook(authorID, bookID uuid.UUID) (*Book, error) {
	var book Book
	err := b.db.Where("author_id = ? AND id = ?", authorID, bookID).First(&book).Error
	return &book, err
}

func (b *bookRepo) GetBooksByPublisher(publisher string) ([]Book, error) {
	var books []Book
	err := b.db.Where("publisher = ?", publisher).Find(&books).Error
	return books, err
}

func (b *bookRepo) DeleteBook(id uuid.UUID) error {
	return b.db.Delete(&Book{}, id).Error
}
