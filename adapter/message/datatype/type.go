package datatype

import (
	"errors"
	"io"

	decode "github.com/valkyriedb/valkyrie/internal/decoder"
)

var (
	ErrUnknownPrimitive = errors.New("unknown primitive type")
	ErrUnknownComposite = errors.New("unknown composite type")
)

type Primitive uint8

const (
	Bool = iota
	Int
	Float
	String
	Blob
)

type Composite uint8

const (
	Prim Composite = iota
	Array
	Map
)

func ParseType(data byte) (Composite, Primitive, error) {
	compByte, primByte := Composite(data>>4), Primitive(data&0b00001111)
	if compByte > Map {
		return 0, 0, ErrUnknownComposite
	}
	if primByte > Blob {
		return 0, 0, ErrUnknownPrimitive
	}

	return compByte, primByte, nil
}

func (p Primitive) Decode(r io.Reader) (any, error) {
	switch p {
	case Bool:
		return decode.Bool(r)
	case Int:
		return decode.Int64(r)
	case Float:
		return decode.Float64(r)
	case String:
		return decode.String(r)
	case Blob:
		return decode.Blob(r)
	default:
		panic(ErrUnknownPrimitive)
	}
}

func (p Primitive) DecodeArray(r io.Reader) (any, error) {
	arrLen, err := decode.Len(r)
	if err != nil {
		return nil, err
	}

	switch p {
	case Bool:
		arr := make([]bool, arrLen)
		for range arrLen {
			v, err := decode.Bool(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		}
		return arr, nil
	case Int:
		arr := make([]int64, arrLen)
		for range arrLen {
			v, err := decode.Int64(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		}
		return arr, nil
	case Float:
		arr := make([]float64, arrLen)
		for range arrLen {
			v, err := decode.Float64(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		}
		return arr, nil
	case String:
		arr := make([]string, arrLen)
		for range arrLen {
			v, err := decode.String(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		}
		return arr, nil
	case Blob:
		arr := make([][]byte, arrLen)
		for range arrLen {
			v, err := decode.Blob(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		}
		return arr, nil
	default:
		panic(ErrUnknownPrimitive)
	}
}
