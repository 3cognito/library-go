package utils

import "errors"

const (
	INVALID_INPUT = "invalid input"
	GENERIC_ERROR = "bad request"
)

var (
	ErrEmailAlreadyExists             = errors.New("email already exists")
	ErrUsernameAlreadyExists          = errors.New("username already exists")
	ErrProfilePictureUrlAlreadyExists = errors.New("profile picture url already exists")
	ErrFileNotProvided                = errors.New("file not provided")
	ErrFileTooLarge                   = errors.New("file too large")
	ErrUnsupportedFileType            = errors.New("unsupported file type")
	ErrISBNAlreadyExists              = errors.New("isbn already exists")
	ErrRecordNotFound                 = errors.New("record not found")
)

const (
	TenMegabytes = 10 * 1024 * 1024
	TwoMegabytes = 2 * 1024 * 1024
)

type FileType string

const (
	Image FileType = "image"
	Book  FileType = "book"
)
