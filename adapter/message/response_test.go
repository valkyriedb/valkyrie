package message

import (
	"bytes"
	"testing"

	"github.com/valkyriedb/valkyrie/adapter/message/status"
)

func TestResponse(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		res := Response{
			Status: status.Unauth,
			Data:   "Unauth",
		}
		if !bytes.Equal(res.ToBytes(), []byte{11, 0, 0, 0, 69, 6, 0, 0, 0, 'U', 'n', 'a', 'u', 't', 'h'}) {
			t.Error("wrong bytes converting")
		}
	})
}
