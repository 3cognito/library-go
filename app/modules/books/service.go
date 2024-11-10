package books

import (
	"github.com/3cognito/library/app/modules/cloudinary"
	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func NewService(
	bookRepo BookRepoInterface,
	cloudinary cloudinary.CloudinaryServiceInterface,
) ServiceInterface {
	return &service{
		bookRepo:   bookRepo,
		cloudinary: cloudinary,
	}
}

func (s *service) AddBook(userId uuid.UUID, data CreateBookRequest) (*Book, error) {
	publicationDate, publicationDateErr := utils.ParseStringTime(data.PublicationDate)
	if publicationDateErr != nil {
		return nil, publicationDateErr
	}
	uploadImageData, uploadImageDataErr := s.cloudinary.UploadFile(data.ImageFile, cloudinary.Image)
	if uploadImageDataErr != nil {
		return nil, uploadImageDataErr
	}

	uploadBookData, uploadBookDataErr := s.cloudinary.UploadFile(data.BookFile, cloudinary.Book)
	if uploadBookDataErr != nil {
		return nil, uploadBookDataErr
	}

	book := &Book{
		Title:              data.Title,
		AuthorID:           userId,
		ISBN:               data.ISBN,
		Publisher:          &data.Publisher,
		PublicationDate:    &publicationDate,
		Pages:              data.Pages,
		Language:           data.Language,
		Description:        &data.Description,
		Genres:             pq.StringArray(data.Genres),
		BookFileUrl:        uploadBookData.URL,
		BookFilePublicID:   uploadBookData.PublicID,
		BookFileSize:       int(uploadBookData.Size),
		BookFileName:       uploadBookData.Name,
		BookFileExtension:  uploadBookData.Extension,
		CoverImageUrl:      uploadImageData.URL,
		CoverImagePublicID: uploadImageData.PublicID,
	}

	if err := s.bookRepo.CreateBook(book); err != nil {
		return nil, err
	}

	return book, nil
}
