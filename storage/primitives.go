package storage

type Primitive interface {
	string | int64 | float64 | bool | []byte
}

type PrimitiveQuery[T Primitive] struct {
	DB  *DB
	key string
}

func (pq *PrimitiveQuery[T]) Get() (T, error) {
	var nothing T
	value, ok := pq.DB.syncMap.Load(pq.key)
	if !ok {
		return nothing, ErrNotFound
	}

	primitive, ok := value.(T)
	if !ok {
		return nothing, ErrWrongType
	}
	return primitive, nil
}

func (pq *PrimitiveQuery[T]) Set(value T) {
	pq.DB.syncMap.Store(pq.key, value)
}

func (pq *PrimitiveQuery[T]) Remove() (T, error) {
	var nothing T
	value, ok := pq.DB.syncMap.LoadAndDelete(pq.key)
	primitive, ok := value.(T)
	if !ok {
		pq.DB.syncMap.Store(pq.key, value)
		return nothing, ErrWrongType
	}
	return primitive, nil
}

func (pq *PrimitiveQuery[T]) Len() (int, error) {
	value, ok := pq.DB.syncMap.Load(pq.key)
	if !ok {
		return 0, ErrNotFound
	}

	switch typed := value.(type) {
	case string:
		return len(typed), nil
	case []byte:
		return len(typed), nil
	default:
		return 0, ErrWrongType
	}
}

func (pq *PrimitiveQuery[T]) Append(postfix string) error {
	value, ok := pq.DB.syncMap.Load(pq.key)
	if !ok {
		return ErrNotFound
	}

	str, ok := value.(string)
	if !ok {
		return ErrWrongType
	}
	pq.DB.syncMap.Store(pq.key, str+postfix)
	return nil
}
func (pq *PrimitiveQuery[T]) Increment() error {
	value, ok := pq.DB.syncMap.Load(pq.key)
	if !ok {
		return ErrNotFound
	}

	num, ok := value.(int64)
	if !ok {
		return ErrWrongType
	}
	pq.DB.syncMap.Store(pq.key, num+1)
	return nil
}

func (pq *PrimitiveQuery[T]) Decrement() error {
	value, ok := pq.DB.syncMap.Load(pq.key)
	if !ok {
		return ErrNotFound
	}

	num, ok := value.(int64)
	if !ok {
		return ErrWrongType
	}
	pq.DB.syncMap.Store(pq.key, num-1)
	return nil
}
