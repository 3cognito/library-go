package bookmarks

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) BookmarkRepoInterface {
	return &bookmarkRepo{
		db: db,
	}
}

func (r *bookmarkRepo) Create(bookmark *Bookmark) error {
	return r.db.Create(bookmark).Error
}

func (r *bookmarkRepo) GetUserBookMarks(userID uuid.UUID) ([]Bookmark, error) {
	var bookmarks []Bookmark

	err := r.db.Joins("Book").Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (r *bookmarkRepo) DeleteBookMark(userID, bookID uuid.UUID) error {
	return r.db.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&Bookmark{}).Error
}
