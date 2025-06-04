package service

import (
	"github.com/valkyriedb/valkyrie/adapter/message"
	"github.com/valkyriedb/valkyrie/adapter/message/datatype"
	"github.com/valkyriedb/valkyrie/adapter/message/op"
	"github.com/valkyriedb/valkyrie/adapter/message/status"
	"github.com/valkyriedb/valkyrie/storage"
)

type Service struct {
	db *storage.DB
}

func New(db *storage.DB) Service {
	return Service{db}
}

func (s Service) Do(req message.Request) message.Response {
	switch req.Composite {
	case datatype.Prim:
		switch req.Primitive {
		case datatype.Bool:
			q := s.db.Bool(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				q.Set(req.Value.(bool))
				return message.NewResponse(nil)
			case op.Pop:
				v, err := q.Remove()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Int:
			q := s.db.Int(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				q.Set(req.Value.(int64))
				return message.NewResponse(nil)
			case op.Pop:
				v, err := q.Remove()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Increment:
				err := q.Increment()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Decrement:
				err := q.Decrement()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			}

		case datatype.Float:
			q := s.db.Float(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				q.Set(req.Value.(float64))
				return message.NewResponse(nil)
			case op.Pop:
				v, err := q.Remove()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.String:
			q := s.db.String(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				q.Set(req.Value.(string))
				return message.NewResponse(nil)
			case op.Pop:
				v, err := q.Remove()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Len:
				l, err := q.Len()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(l)
			case op.Append:
				err := q.Append(req.Value.(string))
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			}

		case datatype.Blob:
			q := s.db.Float(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				q.Set(req.Value.(float64))
				return message.NewResponse(nil)
			case op.Pop:
				v, err := q.Remove()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}
		}

	case datatype.Array:
		switch req.Primitive {
		case datatype.Bool:
			q := s.db.ArrayBool(req.Key)
			switch req.Op {
			case op.Remove:
				v, err := q.Remove(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Insert:
				err := q.Insert(req.Idx, req.Value.([]bool)...)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Len:
				l, err := q.Len()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(l)
			case op.Slice:
				v, err := q.Slice(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Int:
			q := s.db.ArrayInt(req.Key)
			switch req.Op {
			case op.Remove:
				v, err := q.Remove(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Insert:
				err := q.Insert(req.Idx, req.Value.([]int64)...)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Len:
				l, err := q.Len()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(l)
			case op.Slice:
				v, err := q.Slice(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Float:
			q := s.db.ArrayFloat(req.Key)
			switch req.Op {
			case op.Remove:
				v, err := q.Remove(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Insert:
				err := q.Insert(req.Idx, req.Value.([]float64)...)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Len:
				l, err := q.Len()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(l)
			case op.Slice:
				v, err := q.Slice(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.String:
			q := s.db.ArrayString(req.Key)
			switch req.Op {
			case op.Remove:
				v, err := q.Remove(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Insert:
				err := q.Insert(req.Idx, req.Value.([]string)...)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Len:
				l, err := q.Len()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(l)
			case op.Slice:
				v, err := q.Slice(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Blob:
			q := s.db.ArrayBlob(req.Key)
			switch req.Op {
			case op.Remove:
				v, err := q.Remove(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Insert:
				err := q.Insert(req.Idx, req.Value.([][]byte)...)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Len:
				l, err := q.Len()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(l)
			case op.Slice:
				v, err := q.Slice(req.Idx, req.Idx2)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}
		}

	case datatype.Map:
		switch req.Primitive {
		case datatype.Bool:
			q := s.db.MapBool(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				err := q.Set(req.MapKey, req.Value.(bool))
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Remove:
				v, err := q.Remove(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Contains:
				ok, err := q.Contains(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(ok)
			case op.Keys:
				v, err := q.Keys()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Values:
				v, err := q.Values()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Int:
			q := s.db.MapInt(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				err := q.Set(req.MapKey, req.Value.(int64))
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Remove:
				v, err := q.Remove(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Contains:
				ok, err := q.Contains(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(ok)
			case op.Keys:
				v, err := q.Keys()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Values:
				v, err := q.Values()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Float:
			q := s.db.MapFloat(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				err := q.Set(req.MapKey, req.Value.(float64))
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Remove:
				v, err := q.Remove(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Contains:
				ok, err := q.Contains(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(ok)
			case op.Keys:
				v, err := q.Keys()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Values:
				v, err := q.Values()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.String:
			q := s.db.MapString(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				err := q.Set(req.MapKey, req.Value.(string))
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Remove:
				v, err := q.Remove(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Contains:
				ok, err := q.Contains(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(ok)
			case op.Keys:
				v, err := q.Keys()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Values:
				v, err := q.Values()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}

		case datatype.Blob:
			q := s.db.MapBlob(req.Key)
			switch req.Op {
			case op.Get:
				v, err := q.Get(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Set:
				err := q.Set(req.MapKey, req.Value.([]byte))
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(nil)
			case op.Remove:
				v, err := q.Remove(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Contains:
				ok, err := q.Contains(req.MapKey)
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(ok)
			case op.Keys:
				v, err := q.Keys()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			case op.Values:
				v, err := q.Values()
				if err != nil {
					return message.DBErrToRes(err)
				}
				return message.NewResponse(v)
			}
		}
	}

	return message.Response{
		Status: status.UnavailableOp,
	}
}
