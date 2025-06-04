package decode

import (
	"encoding/binary"
	"io"
	"math"
)

func Len(r io.Reader) (int, error) {
	var buf [4]byte
	_, err := io.ReadAtLeast(r, buf[:], 4)
	if err != nil {
		return 0, err
	}

	return int(binary.LittleEndian.Uint32(buf[:])), nil
}

func Bool(r io.Reader) (bool, error) {
	var buf [1]byte
	_, err := io.ReadAtLeast(r, buf[:], 1)
	if err != nil {
		return false, err
	}

	switch buf[0] {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, err
	}
}

func Int64(r io.Reader) (int64, error) {
	var buf [8]byte
	_, err := io.ReadAtLeast(r, buf[:], 8)
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(buf[:])), nil
}

func Float64(r io.Reader) (float64, error) {
	var buf [8]byte
	_, err := io.ReadAtLeast(r, buf[:], 8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(buf[:])), nil
}

func String(r io.Reader) (string, error) {
	data, err := Blob(r)
	return string(data), err
}

func Blob(r io.Reader) ([]byte, error) {
	blobLen, err := Len(r)
	if err != nil {
		return nil, err
	}

	blob := make([]byte, blobLen)
	_, err = io.ReadAtLeast(r, blob, blobLen)
	if err != nil {
		return nil, err
	}

	return blob, nil
}
