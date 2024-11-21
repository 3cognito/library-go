package rating

import "errors"

var (
	ErrRatingOutOfRange = errors.New("rating should be between 1 and 5")
)
