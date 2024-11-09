package books

import "mime/multipart"

type CreateBookRequest struct {
	Title           string                `json:"title"`
	ISBN            string                `json:"isbn"`
	Publisher       string                `json:"publisher"`
	PublicationDate string                `json:"publication_date"`
	Pages           int                   `json:"pages"`
	Language        string                `json:"language"`
	Description     string                `json:"description"`
	Genres          string                `json:"genres"`
	BookFile        *multipart.FileHeader `json:"book_file"`
	ImageFile       *multipart.FileHeader `json:"image_file"`
}
