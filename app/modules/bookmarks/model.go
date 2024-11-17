package bookmarks

import (
	"time"

	"github.com/3cognito/library/app/modules/books"
	"github.com/3cognito/library/app/modules/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bookmark struct {
	ID           string         `json:"id"`
	UserID       uuid.UUID      `gorm:"not null;uniqueIndex:,composite:user_id_book_id" json:"user_id"`
	User         users.User     `json:"-"`
	BookID       uuid.UUID      `gorm:"not null;uniqueIndex:,composite:user_id_book_id" json:"book_id"`
	Book         books.Book     `json:"-"`
	BookmarkedAt *time.Time     `gorm:"type:TIMESTAMP" json:"book_marked_at"`
	CreatedAt    time.Time      `gorm:"not null;type:TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"not null;type:TIMESTAMP;" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"type:TIMESTAMP; index"`
}
