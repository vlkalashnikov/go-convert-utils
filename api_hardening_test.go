package utils

import (
	"testing"
	"time"
)

func TestChunkArrayAlias(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	expected := ChunkArray(input, 2)
	actual := ChunckArray(input, 2)

	if len(expected) != len(actual) {
		t.Fatalf("chunk alias length mismatch: expected %d, got %d", len(expected), len(actual))
	}
	for i := range expected {
		if len(expected[i]) != len(actual[i]) {
			t.Fatalf("chunk %d length mismatch: expected %d, got %d", i, len(expected[i]), len(actual[i]))
		}
		for j := range expected[i] {
			if expected[i][j] != actual[i][j] {
				t.Fatalf("chunk %d index %d mismatch: expected %d, got %d", i, j, expected[i][j], actual[i][j])
			}
		}
	}
}

func TestUniqueBy(t *testing.T) {
	type person struct {
		ID   int
		Name string
	}

	input := []person{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "b"},
		{ID: 1, Name: "c"},
	}

	got := UniqueBy(input, func(p person) int { return p.ID })
	if len(got) != 2 {
		t.Fatalf("expected 2 unique items, got %d", len(got))
	}
	if got[0].Name != "a" || got[1].Name != "b" {
		t.Fatalf("unexpected order/content: %#v", got)
	}
}

func TestUniqSliceHandlesNonComparableKeys(t *testing.T) {
	input := []struct {
		Value string
		Key   []int
	}{
		{Value: "a", Key: []int{1, 2}},
		{Value: "b", Key: []int{1, 2}},
		{Value: "c", Key: []int{3, 4}},
	}

	got := UniqSlice(input, func(v struct {
		Value string
		Key   []int
	}) interface{} {
		return v.Key
	})

	if len(got) != 2 {
		t.Fatalf("expected 2 items for non-comparable keys, got %d", len(got))
	}
	if got[0].Value != "a" || got[1].Value != "c" {
		t.Fatalf("unexpected order/content: %#v", got)
	}
}

func TestIsFirstDayAlias(t *testing.T) {
	first := time.Date(2026, 3, 1, 10, 0, 0, 0, time.UTC)
	notFirst := time.Date(2026, 3, 2, 10, 0, 0, 0, time.UTC)

	if !IsFirstDayOfMonth(first) {
		t.Fatal("expected IsFirstDayOfMonth=true for first day")
	}
	if IsFirstDayOfMonth(notFirst) {
		t.Fatal("expected IsFirstDayOfMonth=false for non-first day")
	}

	if IsFitstDayOfMonth(first) != IsFirstDayOfMonth(first) {
		t.Fatal("deprecated alias differs for first day")
	}
	if IsFitstDayOfMonth(notFirst) != IsFirstDayOfMonth(notFirst) {
		t.Fatal("deprecated alias differs for non-first day")
	}
}
