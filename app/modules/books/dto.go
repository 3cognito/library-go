package books

type CreateBookRequest struct {
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
	Pages           int    `json:"pages"`
	Language        string `json:"language"`
	Description     string `json:"description"`
	Genres          string `json:"genres"`
}
