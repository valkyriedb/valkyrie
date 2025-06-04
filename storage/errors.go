package storage

import "errors"

var (
	ErrWrongType  = errors.New("value of wrong type")
	ErrNotFound   = errors.New("value not found")
	ErrOutOfRange = errors.New("index out of range")
)
