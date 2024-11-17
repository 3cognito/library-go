package books

import (
	"github.com/3cognito/library/app/modules/cloudinary"
	"github.com/3cognito/library/app/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func NewService(
	bookRepo BookRepoInterface,
	deletedBookRepo DeletedBookRepoInterface,
	cloudinary cloudinary.CloudinaryServiceInterface,
) ServiceInterface {
	return &service{
		bookRepo:        bookRepo,
		cloudinary:      cloudinary,
		deletedBookRepo: deletedBookRepo,
	}
}

func (s *service) AddBook(userId uuid.UUID, data CreateBookRequest) (*Book, error) {
	publicationDate, publicationDateErr := utils.ParseStringDate(data.PublicationDate)
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

func (s *service) DeleteBook(userId, bookId uuid.UUID) error {
	book, bookErr := s.bookRepo.GetAuthorBook(userId, bookId)
	if bookErr != nil {
		return ErrResourceNotFound
	}

	if err := s.deleteBook(book); err != nil {
		return err
	}

	go s.deleteBookDataFromCloudinary(book)

	return nil
}

func (s *service) GetAuthorBooks(authorID uuid.UUID) ([]Book, error) {
	return s.bookRepo.GetAuthorBooks(authorID)
}

func (s *service) GetBookByID(id uuid.UUID) (*Book, error) {
	book, err := s.bookRepo.GetBookByID(id)
	if err != nil {
		return nil, ErrResourceNotFound
	}
	return book, nil
}

func (s *service) UpdateBookFiles(authorId, bookId uuid.UUID, data UpdateBookFilesRequest) (*Book, error) {
	book, err := s.bookRepo.GetAuthorBook(authorId, bookId)
	if err != nil {
		return nil, ErrResourceNotFound
	}

	if data.BookFile != nil {
		uploadBookData, uploadBookDataErr := s.cloudinary.UploadFile(data.BookFile, cloudinary.Book)
		if uploadBookDataErr != nil {
			return nil, uploadBookDataErr
		}

		book.BookFileUrl = uploadBookData.URL
		book.BookFilePublicID = uploadBookData.PublicID
		book.BookFileSize = int(uploadBookData.Size)
		book.BookFileName = uploadBookData.Name
		book.BookFileExtension = uploadBookData.Extension
	}

	if data.ImageFile != nil {
		uploadImageData, uploadImageDataErr := s.cloudinary.UploadFile(data.ImageFile, cloudinary.Image)
		if uploadImageDataErr != nil {
			return nil, uploadImageDataErr
		}

		book.CoverImageUrl = uploadImageData.URL
		book.CoverImagePublicID = uploadImageData.PublicID
	}

	if err := s.bookRepo.Save(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *service) UpdateBookDetails(authorId, bookId uuid.UUID, data UpdateBookDetailsRequest) (*Book, error) {
	book, err := s.bookRepo.GetAuthorBook(authorId, bookId)
	if err != nil {
		return nil, ErrResourceNotFound
	}

	publication_date, dateErr := utils.ParseStringDate(data.PublicationDate)
	if dateErr != nil {
		return nil, dateErr
	}

	book.Title = data.Title
	book.ISBN = data.ISBN
	book.Publisher = &data.Publisher
	book.PublicationDate = &publication_date
	book.Pages = data.Pages
	book.Language = data.Language
	book.Description = &data.Description
	book.Genres = pq.StringArray(data.Genres)

	if err := s.bookRepo.Save(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *service) deleteBook(book *Book) error {
	var deletedBook DeletedBook
	tx := s.bookRepo.BeginTrx()
	if err := s.bookRepo.DeleteBook(book.ID); err != nil {
		tx.Rollback()
		return err
	}

	utils.ConvertStruct(book, &deletedBook)

	if err := s.deletedBookRepo.CreateEntry(&deletedBook); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *service) deleteBookDataFromCloudinary(book *Book) error {
	if err := s.cloudinary.DeleteFile(book.BookFilePublicID); err != nil {
		//log error
		return err
	}

	if err := s.cloudinary.DeleteFile(book.CoverImagePublicID); err != nil {
		return err
	}
	return nil
}
