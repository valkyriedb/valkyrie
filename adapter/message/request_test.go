package message

import (
	"bytes"
	"testing"

	"github.com/valkyriedb/valkyrie/adapter/message/datatype"
	"github.com/valkyriedb/valkyrie/adapter/message/op"
)

// | 4byte | byte | byte | 4byte  |  -  |   -    |
// |  Len  | Type |  Op  | KeyLen | Key | Params |

func TestRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		binaryReq := bytes.NewBuffer([]byte{9, 0, 0, 0, 0x14, 2, 3, 0, 0, 0, 'K', 'e', 'y'})

		req, err := ReadRequest(binaryReq)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if req.Composite != datatype.Array {
			t.Errorf("unexpected composite type: got: %d, want: Array(1)", req.Composite)
		}
		if req.Primitive != datatype.Blob {
			t.Errorf("unexpected primitive type: got: %d, want: Blob(4)", req.Primitive)
		}
		if req.Op != op.Pop {
			t.Errorf("unexpected operation: got: %d, want: Pop(2)", req.Op)
		}
		if req.Key != "Key" {
			t.Errorf("unexpected key: got: %s, want: Key", req.Key)
		}
	})
}
