package books

import (
	"time"

	"github.com/3cognito/library/app/modules/users"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model         `json:"-"`
	ID                 uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title              string         `gorm:"not null" json:"title"`
	AuthorID           uuid.UUID      `gorm:"type:uuid;not null;index" json:"author_id"`
	Author             users.User     `json:"-"`
	ISBN               string         `gorm:"unique" json:"isbn"`
	Publisher          *string        `json:"publisher"`
	PublicationDate    *time.Time     `gorm:"type:date" json:"publication_date"`
	Pages              int            `json:"pages"`
	Language           string         `gorm:"not null" json:"language"`
	Description        *string        `json:"description"`
	BookFileUrl        string         `gorm:"not null;unique" json:"book_file_url"`
	BookFilePublicID   string         `gorm:"not null;unique" json:"book_file_public_id"`
	BookFileSize       int            `gorm:"not null" json:"book_file_size"` //size in bytes
	BookFileName       string         `gorm:"not null" json:"book_file_name"`
	BookFileExtension  string         `gorm:"not null" json:"book_file_extension"`
	CoverImageUrl      string         `gorm:"not null;unique" json:"cover_image_url"`
	CoverImagePublicID string         `gorm:"not null;unique" json:"cover_image_public_id"`
	Genres             pq.StringArray `gorm:"type:varchar(64)[]" json:"genres"`
	CreatedAt          time.Time      `gorm:"not null;type:TIMESTAMP;" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"not null;type:TIMESTAMP;" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"type:TIMESTAMP; index"`
}
