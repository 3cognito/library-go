package books

import (
	"github.com/3cognito/library/app/modules/cloudinary"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

type BookRepoInterface interface {
	BeginTrx() *gorm.DB
	CreateBook(book *Book) error
	GetBookByID(id uuid.UUID) (*Book, error)
	Save(book *Book) error
	GetAuthorBooks(authorID uuid.UUID) ([]Book, error)
	GetBooksByPublisher(publisher string) ([]Book, error)
	DeleteBook(id uuid.UUID) error
	GetAuthorBook(authorID, bookID uuid.UUID) (*Book, error)
}

type DeletedBookRepoInterface interface {
	CreateEntry(book *DeletedBook) error
}

type service struct {
	bookRepo        BookRepoInterface
	deletedBookRepo DeletedBookRepoInterface
	cloudinary      cloudinary.CloudinaryServiceInterface
}

type ServiceInterface interface {
	AddBook(userId uuid.UUID, data CreateBookRequest) (*Book, error)
	DeleteBook(userId, bookId uuid.UUID) error
	GetAuthorBooks(authorID uuid.UUID) ([]Book, error)
	// UpdateBook(book *Book) error
	GetBookByID(id uuid.UUID) (*Book, error)
}

type controller struct {
	bookService ServiceInterface
}

type ControllerInterface interface {
	AddBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
	GetAuthorBooks(ctx *gin.Context)
	GetBook(ctx *gin.Context)
}
