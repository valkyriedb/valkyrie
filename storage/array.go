package storage

import "slices"

type ArrayQuery[T Primitive] struct {
	DB  *DB
	key string
}

func (aq *ArrayQuery[T]) Slice(left int, right int) ([]T, error) {
	var nothing []T
	array, err := aq.getArray()
	if err != nil {
		return nothing, err
	}

	if left >= len(array) {
		return nothing, ErrOutOfRange
	}
	if right > len(array) {
		return nothing, ErrOutOfRange
	}

	return array[left:right], nil
}

func (aq *ArrayQuery[T]) Insert(index int, elems ...T) error {
	array, err := aq.getArray()
	if err != nil {
		if err == ErrNotFound {
			array = make([]T, 0)
		} else {
			return err
		}
	}

	if index < 0 && index >= len(array) {
		return ErrOutOfRange
	}

	for _, elem := range elems {
		array[index] = elem
		index++
	}

	aq.DB.syncMap.Store(aq.key, array)
	return nil
}

func (aq *ArrayQuery[T]) Remove(left int, right int) ([]T, error) {
	var nothing []T
	array, err := aq.getArray()
	if err != nil {
		return nothing, err
	}

	if left < 0 || left >= len(array) {
		return nothing, ErrOutOfRange
	}
	if right < 0 || right > len(array) {
		return nothing, ErrOutOfRange
	}

	removed := slices.Delete(array, left, right)
	aq.DB.syncMap.Store(aq.key, removed)
	return removed, nil
}

func (aq *ArrayQuery[T]) Len() (int, error) {
	array, err := aq.getArray()
	if err != nil {
		return 0, err
	}
	return len(array), err
}

func (aq *ArrayQuery[T]) Clear() {
	aq.DB.syncMap.Delete(aq.key)
}

func (aq *ArrayQuery[T]) getArray() ([]T, error) {
	var nothing []T
	value, ok := aq.DB.syncMap.Load(aq.key)
	if !ok {
		return nothing, ErrNotFound
	}

	array, ok := value.([]T)
	if !ok {
		return nothing, ErrWrongType
	}
	return array, nil
}
