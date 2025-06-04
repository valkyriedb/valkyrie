package message

import (
	"bytes"
	"fmt"
	"io"

	"github.com/valkyriedb/valkyrie/adapter/message/datatype"
	"github.com/valkyriedb/valkyrie/adapter/message/op"
	decode "github.com/valkyriedb/valkyrie/internal/decoder"
)

// | 4byte | byte | byte | 4byte  |  -  |   -    |
// |  Len  | Type |  Op  | KeyLen | Key | Params |

type Request struct {
	Composite datatype.Composite
	Primitive datatype.Primitive
	Op        op.Type
	Key       string

	Value  any
	MapKey string
	Idx    int
	Idx2   int
}

func ReadRequest(r io.Reader) (Request, error) {
	l, err := decode.Len(r)
	if err != nil {
		return Request{}, err
	}
	data := make([]byte, l)
	_, err = io.ReadAtLeast(r, data, l)
	if err != nil {
		return Request{}, err
	}

	r = bytes.NewBuffer(data)

	var req Request

	err = req.readHead(r)
	if err != nil {
		return Request{}, fmt.Errorf("can't read head: %w", err)
	}

	err = req.readParams(r)
	if err != nil {
		return Request{}, fmt.Errorf("can't read params: %w", err)
	}

	return req, nil
}

func (req *Request) readHead(r io.Reader) error {
	var head [2]byte
	_, err := io.ReadAtLeast(r, head[:], 2)
	if err != nil {
		return err
	}

	req.Composite, req.Primitive, err = datatype.ParseType(head[0])
	if err != nil {
		return err
	}

	req.Op, err = op.Parse(head[1])
	if err != nil {
		return err
	}

	key, err := decode.String(r)
	if err != nil {
		return err
	}
	req.Key = string(key)

	return nil
}

func (req *Request) readParams(r io.Reader) error {
	switch req.Op {
	case op.Set:
		switch req.Composite {
		case datatype.Prim:
			v, err := req.Primitive.Decode(r)
			if err != nil {
				return err
			}

			req.Value = v
		case datatype.Map:
			key, err := decode.String(r)
			if err != nil {
				return err
			}

			val, err := req.Primitive.Decode(r)
			if err != nil {
				return err
			}

			req.MapKey = key
			req.Value = val
		}
	case op.Append:
		param, err := decode.String(r)
		if err != nil {
			return err
		}

		req.Value = param
	case op.Insert:
		idx, err := decode.Len(r)
		if err != nil {
			return err
		}
		param, err := req.Primitive.DecodeArray(r)
		if err != nil {
			return err
		}

		req.Idx = idx
		req.Value = param
	case op.Remove, op.Slice:
		idxL, err := decode.Len(r)
		if err != nil {
			return err
		}
		idxR, err := decode.Len(r)
		if err != nil {
			return err
		}

		req.Idx = idxL
		req.Idx2 = idxR
	case op.Contains:
		mapKey, err := decode.String(r)
		if err != nil {
			return err
		}

		req.MapKey = mapKey
	}
	return nil
}
