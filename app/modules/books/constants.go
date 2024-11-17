package books

import "errors"

const (
	RequestSuccessful = "request successful"
	RequestFailed     = "request failed"
	BadRequest        = "bad request"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
)
