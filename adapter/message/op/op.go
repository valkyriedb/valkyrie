package op

import "errors"

var (
	ErrUnknownOp = errors.New("unknown op")
)

type Type uint8

const (
	Get Type = iota
	Set
	Pop
	Len
	Append
	Increment
	Decrement
	Insert
	Remove
	Slice
	Contains
	Keys
	Values
)

func Parse(data byte) (Type, error) {
	op := Type(data)
	if op > Values {
		return 0, ErrUnknownOp
	}

	return op, nil
}
