package cloudinary

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func NewService(
	config config.Config,
) CloudinaryServiceInterface {
	client, _ := cloudinary.NewFromURL(config.CloudinaryURL)
	return &cloudinaryService{
		client: client,
	}

}

func (c *cloudinaryService) UploadFile(file *multipart.FileHeader, fileType FileType) (FileData, error) {
	var data FileData
	fileData, validationErr := utils.ValidateFile(file, c.parseFileType(fileType))
	if validationErr != nil {
		return data, validationErr
	}

	openedFile, openErr := file.Open()
	if openErr != nil {
		return data, openErr
	}
	defer openedFile.Close()

	uploadParams := uploader.UploadParams{
		Folder: CLOUDINARY__APP_FOLDER,
	}

	ctx := context.Background()

	uploadResult, uploadErr := c.client.Upload.Upload(ctx, openedFile, uploadParams)
	if uploadErr != nil {
		return data, uploadErr
	}

	if uploadResult.Error.Message != "" {
		return data, errors.New(uploadResult.Error.Message)
	}

	data.Name = fileData.Name
	data.Size = fileData.Size
	data.URL = uploadResult.SecureURL
	data.Extension = uploadResult.Format
	data.PublicID = uploadResult.PublicID

	return data, nil
}

func (c *cloudinaryService) DeleteFile(publicID string) error {
	ctx := context.Background()
	_, err := c.client.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *cloudinaryService) parseFileType(fileType FileType) utils.FileType {
	switch fileType {
	case Image:
		return "image"
	case Book:
		return "book"
	default:
		return ""
	}
}
