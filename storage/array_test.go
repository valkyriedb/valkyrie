package storage

import (
	"slices"
	"testing"
)

func TestInsertAndSlice(t *testing.T) {
	db := DB{}
	k, v := "k", "v"
	s := []string{v, v, v}
	if err := db.ArrayString(k).Insert(0, s...); err != nil {
		t.Fatalf("could not insert value: %v", err)
	}

	got, err := db.ArrayString(k).Slice(0, 3)
	if err != nil {
		t.Fatalf("could not get slice: %v", err)
	}

	if !slices.Equal(got, s) {
		t.Errorf("got %v, expected %v", got, s)
	}

	if err := db.ArrayString(k).Insert(-1, s...); err == nil {
		t.Fatalf("expected error: %v", ErrOutOfRange)
	}

	if _, err := db.ArrayString(k).Slice(-1, -1); err == nil {
		t.Fatalf("expected error: %v", ErrOutOfRange)
	}
}

func TestArrayRemove(t *testing.T) {
	db := DB{}
	k, v := "k", "v"
	s := []string{v, v, v}
	if err := db.ArrayString(k).Insert(0, s...); err != nil {
		t.Fatalf("could not insert value: %v", err)
	}

	got, err := db.ArrayString(k).Remove(0, 3)
	if err != nil {
		t.Errorf("could not remove value: %v", err)
	}

	if !slices.Equal(got, s) {
		t.Errorf("got %v, expected %v", got, s)
	}

	if _, err := db.ArrayString(k).Slice(0, 3); err == nil {
		t.Fatalf("expected error: %v", ErrOutOfRange)
	}

	if _, err := db.ArrayString(k).Remove(-1, -1); err == nil {
		t.Fatalf("expected error: %v", ErrOutOfRange)
	}
}

func TestArrayLen(t *testing.T) {
	db := DB{}
	k, v := "k", "v"
	s := []string{v, v, v}
	if err := db.ArrayString(k).Insert(0, s...); err != nil {
		t.Fatalf("could not insert value: %v", err)
	}

	l, err := db.ArrayString(k).Len()
	if err != nil {
		t.Fatalf("could not get length: %v", err)
	}

	expected := len(s)
	if l != expected {
		t.Errorf("got length %d, expected %d", l, expected)
	}
}

func TestClear(t *testing.T) {
	db := DB{}
	k, v := "k", "v"
	s := []string{v, v, v}
	if err := db.ArrayString(k).Insert(0, s...); err != nil {
		t.Fatalf("could not insert value: %v", err)
	}
	db.ArrayString(k).Clear()
	if _, err := db.ArrayString(k).Len(); err == nil {
		t.Errorf("expected error: %v", ErrNotFound)
	}
}
