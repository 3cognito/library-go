package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

type FileData struct {
	Name      string
	Extension string
	Size      int64
	URL       string
}

func ValidateFile(file *multipart.FileHeader) (FileData, error) {
	var data FileData
	if file == nil {
		return data, ErrFileNotProvided
	}

	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	if !isSupportedFileExtension(fileExtension) {
		return data, ErrUnsupportedFileType
	}

	if file.Size > TenMegabytes {
		return data, ErrFileTooLarge
	}

	return FileData{
		Name:      file.Filename,
		Extension: fileExtension,
		Size:      file.Size,
	}, nil
}

func isSupportedFileExtension(fileName string) bool {
	supportedExtensions := []string{"pdf", "epub"}
	for _, ext := range supportedExtensions {
		if ext == fileName {
			return true
		}
	}

	return false
}
