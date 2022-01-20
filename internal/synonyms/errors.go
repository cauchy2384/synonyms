package synonyms

import "errors"

var (
	ErrValiadation = errors.New("validation error")
	ErrNotFound    = errors.New("not found")
)
