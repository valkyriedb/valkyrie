package encode

import (
	"encoding/binary"
	"math"
	"slices"
)

func AppendAny(b []byte, a any) []byte {
	switch vv := a.(type) {
	case bool:
		return appendBool(b, vv)
	case int64:
		return appendInt64(b, vv)
	case float64:
		return appendFloat64(b, vv)
	case string:
		return appendString(slices.Grow(b, len(vv)+4), vv)
	case []byte:
		return appendBlob(slices.Grow(b, len(vv)+4), vv)
	case []bool:
		b := slices.Grow(b, len(vv))
		for _, v := range vv {
			b = appendBool(b, v)
		}
		return b
	case []int64:
		b := slices.Grow(b, len(vv)*8)
		for _, v := range vv {
			b = appendInt64(b, v)
		}
		return b
	case []float64:
		b := slices.Grow(b, len(vv)*8)
		for _, v := range vv {
			b = appendFloat64(b, v)
		}
		return b
	case []string:
		var totalLen int
		for _, v := range vv {
			totalLen += len(v) + 4
		}
		b := slices.Grow(b, totalLen)
		for _, v := range vv {
			b = appendString(b, v)
		}
		return b
	case [][]byte:
		var totalLen int
		for _, v := range vv {
			totalLen += len(v) + 4
		}
		b := slices.Grow(b, totalLen)
		for _, v := range vv {
			b = appendBlob(b, v)
		}
		return b
	default:
		panic("unsupported type")
	}
}

func PutLen(b []byte, len int) {
	binary.LittleEndian.PutUint32(b, uint32(len))
}

func appendLen(b []byte, len int) []byte {
	return binary.LittleEndian.AppendUint32(b, uint32(len))
}

func appendBool(b []byte, bl bool) []byte {
	var data byte
	if bl {
		data = 1
	}
	return append(b, data)
}

func appendInt64(b []byte, i int64) []byte {
	return binary.LittleEndian.AppendUint64(b, uint64(i))
}

func appendFloat64(b []byte, f float64) []byte {
	return binary.LittleEndian.AppendUint64(b, math.Float64bits(f))
}

func appendString(b []byte, s string) []byte {
	return appendBlob(b, []byte(s))
}

func appendBlob(b []byte, blob []byte) []byte {
	b = appendLen(b, len(blob))
	b = append(b, blob...)
	return b
}
