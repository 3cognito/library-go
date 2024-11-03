package utils

import "strings"

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

	return err
}
