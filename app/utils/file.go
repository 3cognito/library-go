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

func ValidateFile(file *multipart.FileHeader, fileType FileType) (FileData, error) {
	var data FileData
	if file == nil {
		return data, ErrFileNotProvided
	}

	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	isSupported := (fileType == "image" && isSupportedImageExtension(fileExtension)) ||
		(fileType == "book" && isSupportedBookExtension(fileExtension))

	if !isSupported {
		return data, ErrUnsupportedFileType
	}

	isSizeValid := (fileType == "image" && file.Size <= TwoMegabytes) || (fileType == "book" && file.Size <= TenMegabytes)
	if !isSizeValid {
		return data, ErrFileTooLarge
	}

	return FileData{
		Name:      file.Filename,
		Extension: fileExtension,
		Size:      file.Size,
	}, nil
}

func isSupportedImageExtension(fileName string) bool {
	supportedExtensions := []string{"jpg", "jpeg", "png"}
	for _, ext := range supportedExtensions {
		if ext == fileName {
			return true
		}
	}

	return false
}

func isSupportedBookExtension(fileName string) bool {
	supportedExtensions := []string{"pdf", "epub"}
	for _, ext := range supportedExtensions {
		if ext == fileName {
			return true
		}
	}

	return false
}
