package base

import (
	"github.com/3cognito/library/app/modules/auth"
	"github.com/3cognito/library/app/modules/bookmarks"
	"github.com/3cognito/library/app/modules/books"
)

func (b *base) WithAuthController() auth.AuthControllerInterface {
	return auth.NewAuthController(b.WithAuthService())
}

func (b *base) WithBookController() books.ControllerInterface {
	return books.NewController(b.WithBookService())
}

func (b *base) WithBookmarkController() bookmarks.BookmarkControllerInterface {
	return bookmarks.NewController(b.WithBookmarkService())
}
