package storage

import "fmt"

var (
	ErrWrongType  = fmt.Errorf("value of wrong type")
	ErrNotFound   = fmt.Errorf("value not found")
	ErrOutOfRange = fmt.Errorf("index out of range")
)
