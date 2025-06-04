package storage

type MapQuery[T Primitive] struct {
	DB  *DB
	key string
}

func (mq *MapQuery[T]) Get(key string) (T, error) {
	var nothing T
	m, err := mq.getMap()
	if err != nil {
		return nothing, err
	}

	value, ok := m[key]
	if !ok {
		return nothing, ErrNotFound
	}
	return value, nil
}

func (mq *MapQuery[T]) Set(key string, value T) error {
	m, err := mq.getMap()
	if err != nil {
		return err
	}

	m[key] = value
	return err
}

func (mq *MapQuery[T]) Remove(key string) (T, error) {
	var nothing T
	m, err := mq.getMap()
	if err != nil {
		return nothing, err
	}

	value, ok := m[key]
	if !ok {
		return nothing, ErrNotFound
	}
	delete(m, key)
	return value, nil
}

func (mq *MapQuery[T]) Contains(key string) (bool, error) {
	m, err := mq.getMap()
	if err != nil {
		return false, err
	}

	_, ok := m[key]
	if !ok {
		return false, nil
	}
	return true, nil
}

func (mq *MapQuery[T]) Len() (int, error) {
	m, err := mq.getMap()
	if err != nil {
		return 0, err
	}
	return len(m), err
}

func (mq *MapQuery[T]) Keys() ([]string, error) {
	var nothing []string
	m, err := mq.getMap()
	if err != nil {
		return nothing, err
	}

	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys, err
}

func (mq *MapQuery[T]) Values() ([]T, error) {
	var nothing []T
	m, err := mq.getMap()
	if err != nil {
		return nothing, err
	}

	values := make([]T, len(m))
	var i int
	for _, v := range m {
		values[i] = v
		i++
	}
	return values, err
}

func (mq *MapQuery[T]) getMap() (map[string]T, error) {
	var nothing map[string]T
	value, ok := mq.DB.syncMap.Load(mq.key)
	if !ok {
		return nothing, ErrNotFound
	}

	m, ok := value.(map[string]T)
	if !ok {
		return nothing, ErrWrongType
	}
	return m, nil
}
