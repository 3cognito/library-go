package utils

import (
	"strings"

	commons "github.com/3cognito/library/app/common"
)

func CheckUniqueConstrainstErr(err error) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	if strings.Contains(errMsg, "email") {
		return ErrEmailAlreadyExists
	}

	if strings.Contains(errMsg, "username") {
		return ErrUsernameAlreadyExists
	}

	if strings.Contains(errMsg, "profile_picture_url") {
		return ErrProfilePictureUrlAlreadyExists
	}

	if strings.Contains(errMsg, "isbn") {
		return ErrISBNAlreadyExists
	}

	if strings.Contains(errMsg, "idx_bookmarks_user_id_book_id") {
		return commons.ErrBookAlreadyBookmarked
	}

	return err
}
