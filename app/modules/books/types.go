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
	CreateBook(book *Book) error
	GetBookByID(id uuid.UUID) (*Book, error)
	Save(book *Book) error
	GetAuthorBooks(authorID uuid.UUID) ([]Book, error)
	GetBooksByPublisher(publisher string) ([]Book, error)
	DeleteBook(id uuid.UUID) error
	GetAuthorBook(authorID, bookID uuid.UUID) (*Book, error)
}

type service struct {
	bookRepo   BookRepoInterface
	cloudinary cloudinary.CloudinaryServiceInterface
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
