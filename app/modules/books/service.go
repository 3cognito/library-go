package books

import "github.com/google/uuid"

func NewService(
	bookRepo BookRepoInterface,
) ServiceInterface {
	return &service{
		bookRepo: bookRepo,
	}
}

func (s *service) CreateBook(userId uuid.UUID, data CreateBookRequest) (*Book, error) {

	return nil, nil
}
