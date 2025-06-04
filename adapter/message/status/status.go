package status

type Type uint8

const (
	OK            Type = 0
	InvalidReq    Type = 64
	UnavailableOp Type = 65
	Unauth        Type = 69
	NotFound      Type = 128
	WrongType     Type = 129
	OutOfRange    Type = 130
	InternalErr   Type = 255
)

func (status Type) ToByte() byte {
	return byte(status)
}
