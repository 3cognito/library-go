package cloudinary

import (
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
)

type cloudinaryService struct {
	client *cloudinary.Cloudinary
}

type CloudinaryServiceInterface interface {
	UploadFile(file *multipart.FileHeader, fileType FileType) (FileData, error)
	DeleteFile(publicID string) error
}

type FileData struct {
	Name      string
	Extension string
	Size      int64
	URL       string
	PublicID  string
}
