package message

import (
	"errors"

	"github.com/valkyriedb/valkyrie/adapter/message/status"
	encode "github.com/valkyriedb/valkyrie/internal/encoder"
	"github.com/valkyriedb/valkyrie/storage"
)

// | 4byte |  byte  |  -   |
// |  Len  | Status | Data |

type Response struct {
	Status status.Type
	Data   any
}

func (res Response) ToBytes() []byte {
	data := make([]byte, 5)
	if res.Data != nil {
		data = encode.AppendAny(data, res.Data)
	}
	encode.PutLen(data, len(data[4:]))
	data[4] = res.Status.ToByte()
	return data
}

func DBErrToRes(err error) Response {
	if errors.Is(err, storage.ErrNotFound) {
		return Response{Status: status.NotFound}
	} else if errors.Is(err, storage.ErrWrongType) {
		return Response{Status: status.WrongType}
	} else if errors.Is(err, storage.ErrOutOfRange) {
		return Response{Status: status.OutOfRange}
	}
	panic("unknown db error")
}

func NewResponse(data any) Response {
	return Response{
		Status: status.OK,
		Data:   data,
	}
}
