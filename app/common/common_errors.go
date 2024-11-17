package commons

import "errors"

var (
	ErrResourceNotFound      = errors.New("resource not found")
	ErrBookAlreadyBookmarked = errors.New("book already bookmarked")
)
