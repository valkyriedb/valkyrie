package message

import (
	"github.com/valkyriedb/valkyrie/adapter/message/status"
	encode "github.com/valkyriedb/valkyrie/internal/encoder"
)

// | 4byte |  byte  |  -   |
// |  Len  | Status | Data |

type Response struct {
	Status status.Type
	Data   any
}

func (res Response) ToBytes() []byte {
	data := make([]byte, 5)
	data = encode.AppendAny(data, res.Data)
	encode.PutLen(data, len(data[4:]))
	data[4] = res.Status.ToByte()
	return data
}
