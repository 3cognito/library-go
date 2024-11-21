package utils

import (
	"strings"

	commons "github.com/3cognito/library/app/common"
	"gorm.io/gorm"
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

func FindRecordErr(res *gorm.DB) error {
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
