package storage

import (
	"slices"
	"testing"
)

func TestMapSetAndGet(t *testing.T) {
	var db DB
	mk, k, v := "mk", "k", "v"
	if err := db.MapString(mk).Set(k, v); err != nil {
		t.Fatalf("could not set value to key: %v", err)
	}

	got, err := db.MapString(mk).Get(k)
	if err != nil {
		t.Fatalf("could not get value from key: %v", err)
	}
	if got != v {
		t.Errorf("got %v, expected %v", got, v)
	}
}

func TestMapRemove(t *testing.T) {
	var db DB
	mk, k, v := "mk", "k", "v"
	if err := db.MapString(mk).Set(k, v); err != nil {
		t.Fatalf("could not set value to key: %v", err)
	}

	got, err := db.MapString(mk).Remove(k)
	if err != nil {
		t.Fatalf("could not remove value from key: %v", err)
	}
	if got != v {
		t.Errorf("got %v, expected %v", got, v)
	}
	if _, err := db.MapString(mk).Get(k); err == nil {
		t.Errorf("expected error: %v", ErrNotFound)
	}
}

func TestContains(t *testing.T) {
	var db DB
	mk, k, v := "mk", "k", "v"
	if err := db.MapString(mk).Set(k, v); err != nil {
		t.Fatalf("could not set value to key: %v", err)
	}

	got, err := db.MapString(mk).Contains(k)
	if err != nil {
		t.Fatalf("could not check key: %v", err)
	}
	if got != true {
		t.Errorf("expected to contain key")
	}
}

func TestMapLen(t *testing.T) {
	var db DB
	mk := "mk"
	m := map[string]string{"k1": "v1", "k2": "v2"}
	for k, v := range m {
		if err := db.MapString(mk).Set(k, v); err != nil {
			t.Fatalf("could not set value to key: %v", err)
		}
	}
	l, err := db.MapString(mk).Len()
	if err != nil {
		t.Fatalf("could not get map length: %v", err)
	}
	expected := len(m)
	if l != expected {
		t.Errorf("got length %d, expected %d", l, expected)
	}
}

func TestKeys(t *testing.T) {
	var db DB
	mk := "mk"
	m := map[string]string{"k1": "v1", "k2": "v2"}
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	for k, v := range m {
		if err := db.MapString(mk).Set(k, v); err != nil {
			t.Fatalf("could not set value to key: %v", err)
		}
	}

	got, err := db.MapString(mk).Keys()
	if err != nil {
		t.Fatalf("could not get keys: %v", err)
	}
	if slices.Equal(keys, got) {
		t.Errorf("got %v, expected %v", got, keys)
	}
}

func TestValues(t *testing.T) {
	var db DB
	mk := "mk"
	m := map[string]string{"k1": "v1", "k2": "v2"}
	values := make([]string, len(m))
	var i int
	for _, v := range m {
		values[i] = v
		i++
	}
	for k, v := range m {
		if err := db.MapString(mk).Set(k, v); err != nil {
			t.Fatalf("could not set value to key: %v", err)
		}
	}

	got, err := db.MapString(mk).Values()
	if err != nil {
		t.Fatalf("could not get keys: %v", err)
	}
	if slices.Equal(values, got) {
		t.Errorf("got %v, expected %v", got, values)
	}
}
