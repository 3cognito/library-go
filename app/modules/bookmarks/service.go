package bookmarks

import (
	commons "github.com/3cognito/library/app/common"
	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
)

func NewService(
	repo BookmarkRepoInterface,
) BookmarkServiceInterface {
	return &bookmarkService{
		repo: repo,
	}
}

//TODO: SEARCH USER BOOKMARKS

func (s *bookmarkService) AddToBookmark(userId, bookId uuid.UUID) error {
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
