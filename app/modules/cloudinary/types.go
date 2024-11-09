package cloudinary

import (
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
)

type cloudinaryService struct {
	client *cloudinary.Cloudinary
}

type CloudinaryServiceInterface interface {
	UploadImage(file *multipart.FileHeader, fileType FileType) (FileData, error)
	DeleteImage(publicID string) error
}

type FileData struct {
	Name      string
	Extension string
	Size      int64
	URL       string
	PublicID  string
}
