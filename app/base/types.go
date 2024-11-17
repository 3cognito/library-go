package base

import (
	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/modules/auth"
	"github.com/3cognito/library/app/modules/bookmarks"
	"github.com/3cognito/library/app/modules/books"
	"gorm.io/gorm"
)

type base struct {
	configs config.Config
	db      *gorm.DB
}

type appControllers struct {
	AuthC      auth.AuthControllerInterface
	BooksC     books.ControllerInterface
	BookmarksC bookmarks.BookmarkControllerInterface
}
