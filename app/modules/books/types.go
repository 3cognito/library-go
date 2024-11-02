package books

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

type BookRepoInterface interface {
	CreateBook(book *Book) error
	GetBookByID(id uint) (*Book, error)
	Save(book *Book) error
	GetAuthorBooks(authorID uuid.UUID) ([]Book, error)
	GetBooksByPublisher(publisher string) ([]Book, error)
}
