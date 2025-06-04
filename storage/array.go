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

	if valid := isValidRange(left, right, len(array)); !valid {
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

	if index < 0 || index > len(array) {
		return ErrOutOfRange
	}

	array = slices.Insert(array, index, elems...)
	aq.DB.syncMap.Store(aq.key, array)
	return nil
}

func (aq *ArrayQuery[T]) Remove(left int, right int) ([]T, error) {
	var nothing []T
	array, err := aq.getArray()
	if err != nil {
		return nothing, err
	}

	if valid := isValidRange(left, right, len(array)); !valid {
		return nothing, ErrOutOfRange
	}

	removed := make([]T, right-left)
	copy(removed, array[left:right])
	aq.DB.syncMap.Store(aq.key, slices.Delete(array, left, right))
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

func isValidRange(left int, right int, length int) bool {
	if left < 0 || left > length {
		return false
	}
	if right < 0 || right > length {
		return false
	}
	if left > right {
		return false
	}
	return true
}
