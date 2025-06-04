package storage

import (
	"testing"
)

func TestSetAndGet(t *testing.T) {
	db := DB{}
	k, v := "k", "v"
	db.String(k).Set(v)
	got, err := db.String(k).Get()
	if err != nil {
		t.Fatalf("could not get value: %v", err)
	}
	if got != v {
		t.Errorf("got %v, expected %v", got, v)
	}
}

func TestRemove(t *testing.T) {
	db := DB{}
	k, v := "k", "v"
	db.String(k).Set(v)
	got, err := db.String(k).Remove()
	if err != nil {
		t.Fatalf("could not delete value: %v", err)
	}
	if got != v {
		t.Errorf("got %v, expected %v", got, v)
	}
	got, err = db.String(k).Get()
	if got != "" {
		t.Errorf("expected empty value, got %v", got)
	}
	if err != ErrNotFound {
		t.Errorf("expected error: %v", ErrNotFound)
	}
}

func TestLen(t *testing.T) {
	db := DB{}
	k, v := "k", "abcde"
	db.String(k).Set(v)
	l, err := db.String(k).Len()
	if err != nil {
		t.Fatalf("could not get length: %v", err)
	}
	expected := len(v)
	if l != expected {
		t.Errorf("got length %d, expected %d", l, expected)
	}
}

func TestAppend(t *testing.T) {
	db := DB{}
	k, v := "k", "abcd"
	postfix := "efgh"
	db.String(k).Set(v)
	if err := db.String(k).Append(postfix); err != nil {
		t.Fatalf("could not append string: %v", err)
	}
	got, err := db.String(k).Get()
	if err != nil {
		t.Fatalf("could not get value: %v", err)
	}
	expected := v + postfix
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestIncrement(t *testing.T) {
	db := DB{}
	k, v := "k", int64(1)
	db.Int(k).Set(v)
	if err := db.Int(k).Increment(); err != nil {
		t.Fatalf("could not increment number: %v", err)
	}
	got, err := db.Int(k).Get()
	if err != nil {
		t.Fatalf("could not get value: %v", err)
	}
	expected := v + 1
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestDecrement(t *testing.T) {
	db := DB{}
	k, v := "k", int64(1)
	db.Int(k).Set(v)
	if err := db.Int(k).Decrement(); err != nil {
		t.Fatalf("could not increment number: %v", err)
	}
	got, err := db.Int(k).Get()
	if err != nil {
		t.Fatalf("could not get value: %v", err)
	}
	expected := v - 1
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
