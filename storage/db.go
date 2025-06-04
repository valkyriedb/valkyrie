package storage

import (
	"sync"
)

type DB struct {
	syncMap sync.Map
}

func (db *DB) String(key string) *PrimitiveQuery[string] {
	pq := &PrimitiveQuery[string]{
		DB:  db,
		key: key,
	}
	return pq
}

func (db *DB) Int(key string) *PrimitiveQuery[int64] {
	pq := &PrimitiveQuery[int64]{
		DB:  db,
		key: key,
	}
	return pq
}

func (db *DB) Float(key string) *PrimitiveQuery[float64] {
	pq := &PrimitiveQuery[float64]{
		DB:  db,
		key: key,
	}
	return pq
}

func (db *DB) Bool(key string) *PrimitiveQuery[bool] {
	pq := &PrimitiveQuery[bool]{
		DB:  db,
		key: key,
	}
	return pq
}

func (db *DB) Blob(key string) *PrimitiveQuery[[]byte] {
	pq := &PrimitiveQuery[[]byte]{
		DB:  db,
		key: key,
	}
	return pq
}

func (db *DB) ArrayString(key string) *ArrayQuery[string] {
	aq := &ArrayQuery[string]{
		DB:  db,
		key: key,
	}
	return aq
}

func (db *DB) ArrayInt(key string) *ArrayQuery[int64] {
	aq := &ArrayQuery[int64]{
		DB:  db,
		key: key,
	}
	return aq
}

func (db *DB) ArrayFloat(key string) *ArrayQuery[float64] {
	aq := &ArrayQuery[float64]{
		DB:  db,
		key: key,
	}
	return aq
}

func (db *DB) ArrayBool(key string) *ArrayQuery[bool] {
	aq := &ArrayQuery[bool]{
		DB:  db,
		key: key,
	}
	return aq
}

func (db *DB) ArrayBlob(key string) *ArrayQuery[[]byte] {
	aq := &ArrayQuery[[]byte]{
		DB:  db,
		key: key,
	}
	return aq
}

func (db *DB) MapString(key string) *MapQuery[string] {
	mq := &MapQuery[string]{
		DB:  db,
		key: key,
	}
	return mq
}

func (db *DB) MapInt(key string) *MapQuery[int64] {
	mq := &MapQuery[int64]{
		DB:  db,
		key: key,
	}
	return mq
}

func (db *DB) MapFloat(key string) *MapQuery[float64] {
	mq := &MapQuery[float64]{
		DB:  db,
		key: key,
	}
	return mq
}

func (db *DB) MapBool(key string) *MapQuery[bool] {
	mq := &MapQuery[bool]{
		DB:  db,
		key: key,
	}
	return mq
}

func (db *DB) MapBlob(key string) *MapQuery[[]byte] {
	mq := &MapQuery[[]byte]{
		DB:  db,
		key: key,
	}
	return mq
}
