package bookmarks

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookmarkRepo struct {
	db *gorm.DB
}

type bookmarkService struct {
	repo BookmarkRepoInterface
}

type bookmarkController struct {
	service BookmarkServiceInterface
}

type BookmarkRepoInterface interface {
	Create(bookmark *Bookmark) error
	GetUserBookMarks(userID uuid.UUID) ([]Bookmark, error)
	DeleteBookMark(userID, bookID uuid.UUID) error
}

type BookmarkServiceInterface interface {
	AddToBookmark(userId, bookId uuid.UUID) error
	RemoveFromBookmark(userId, bookId uuid.UUID) error
	GetUserBookMarks(userId uuid.UUID) ([]Bookmark, error)
}

type BookmarkControllerInterface interface {
	AddToBookmark(ctx *gin.Context)
	RemoveFromBookmark(ctx *gin.Context)
	GetUserBookMarks(ctx *gin.Context)
}
