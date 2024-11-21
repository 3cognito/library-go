package bookmarks

import (
	commons "github.com/3cognito/library/app/common"
	"github.com/3cognito/library/app/modules/books"
	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
)

func NewService(
	repo BookmarkRepoInterface,
	bookRepo books.BookRepoInterface,
) BookmarkServiceInterface {
	return &bookmarkService{
		repo:     repo,
		bookRepo: bookRepo,
	}
}

//TODO: SEARCH USER BOOKMARKS

func (s *bookmarkService) AddToBookmark(userId, bookId uuid.UUID) error {
	_, bookErr := s.bookRepo.GetBookByID(bookId)
	if bookErr != nil {
		return commons.ErrResourceNotFound
	}

	now := utils.TimeNow()
	bookmark := &Bookmark{
		UserID:       userId,
		BookID:       bookId,
		BookmarkedAt: &now,
	}

	if err := s.repo.Create(bookmark); err != nil {
		return err
	}

	return nil
}

func (s *bookmarkService) RemoveFromBookmark(userId, bookId uuid.UUID) error {
	err := s.repo.DeleteBookMark(userId, bookId)
	if err != nil {
		return commons.ErrResourceNotFound
	}

	return nil
}

func (s *bookmarkService) GetUserBookMarks(userId uuid.UUID) ([]Bookmark, error) {
	return s.repo.GetUserBookMarks(userId)
}
