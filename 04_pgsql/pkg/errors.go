package pkg

import "errors"

var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("object exists in the DB")
	ErrNotExists     = errors.New("object not exists in the DB")
)
