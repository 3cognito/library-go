package bookmarks

import (
	"time"

	"github.com/3cognito/library/app/modules/books"
	"github.com/3cognito/library/app/modules/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookMark struct {
	ID           string         `json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null,index:,composite:idx_user_book" json:"user_id"`
	User         users.User     `json:"-"`
	BookID       uuid.UUID      `gorm:"type:uuid;not null,index:,composite:idx_user_book" json:"book_id"`
	Book         books.Book     `json:"-"`
	BookMarkedAt *time.Time     `gorm:"type:TIMESTAMP" json:"book_marked_at"`
	CreatedAt    time.Time      `gorm:"not null;type:TIMESTAMP;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"not null;type:TIMESTAMP;" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"type:TIMESTAMP; index"`
}
