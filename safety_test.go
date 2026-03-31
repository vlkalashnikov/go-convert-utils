package utils

import (
	"fmt"
	"math"
	"testing"
)

func TestChunckArrayInvalidChunkSize(t *testing.T) {
	cases := []int{0, -1}
	for _, chunkSize := range cases {
		t.Run(fmt.Sprintf("chunkSize_%d", chunkSize), func(t *testing.T) {
			got := ChunckArray([]int{1, 2, 3}, chunkSize)
			if got != nil {
				t.Fatalf("expected nil for invalid chunk size %d, got %#v", chunkSize, got)
			}
		})
	}
}

func TestChunckArrayValid(t *testing.T) {
	got := ChunckArray([]int{1, 2, 3, 4, 5}, 2)
	if len(got) != 3 {
		t.Fatalf("expected 3 chunks, got %d", len(got))
	}

	expected := [][]int{{1, 2}, {3, 4}, {5}}
	for i := range expected {
		if len(got[i]) != len(expected[i]) {
			t.Fatalf("chunk %d length mismatch: expected %d, got %d", i, len(expected[i]), len(got[i]))
		}
		for j := range expected[i] {
			if got[i][j] != expected[i][j] {
				t.Fatalf("chunk %d index %d mismatch: expected %d, got %d", i, j, expected[i][j], got[i][j])
			}
		}
	}
}

func TestB2F64ShortInput(t *testing.T) {
	if got := B2F64([]byte{1, 2, 3}); got != 0 {
		t.Fatalf("expected 0 for short input, got %v", got)
	}
}

func TestB2F64Roundtrip(t *testing.T) {
	const expected = 12.34
	got := B2F64(F642B(expected))
	if math.Abs(got-expected) > 1e-9 {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestRmElemByIdxOutOfRange(t *testing.T) {
	original := []int{10, 20, 30}

	got := RmElemByIdx(original, -1)
	if len(got) != len(original) {
		t.Fatalf("negative index must keep slice unchanged")
	}
	for i := range original {
		if got[i] != original[i] {
			t.Fatalf("negative index changed slice at %d", i)
		}
	}

	got = RmElemByIdx(original, len(original))
	if len(got) != len(original) {
		t.Fatalf("overflow index must keep slice unchanged")
	}
	for i := range original {
		if got[i] != original[i] {
			t.Fatalf("overflow index changed slice at %d", i)
		}
	}
}

func TestRmElemByIdxValid(t *testing.T) {
	got := RmElemByIdx([]int{10, 20, 30}, 1)
	expected := []int{10, 30}

	if len(got) != len(expected) {
		t.Fatalf("expected len %d, got %d", len(expected), len(got))
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("index %d mismatch: expected %d, got %d", i, expected[i], got[i])
		}
	}
}

func TestFnNameInvalidInput(t *testing.T) {
	cases := []interface{}{nil, 123, "not-fn"}
	for _, tc := range cases {
		if got := FnName(tc); got != "" {
			t.Fatalf("expected empty function name for %T, got %q", tc, got)
		}
	}
}

func TestFnNameValid(t *testing.T) {
	got := FnName(TestFnNameValid)
	if got == "" {
		t.Fatal("expected function name for function input")
	}
}

func TestJsonBScanTypes(t *testing.T) {
	var j JsonB

	if err := j.Scan([]byte(`{"x":1}`)); err != nil {
		t.Fatalf("scan []byte failed: %v", err)
	}
	if j["x"] != float64(1) {
		t.Fatalf("expected x=1, got %#v", j["x"])
	}

	if err := j.Scan(`{"y":2}`); err != nil {
		t.Fatalf("scan string failed: %v", err)
	}
	if _, ok := j["x"]; ok {
		t.Fatal("expected map replacement, key x must not persist")
	}
	if j["y"] != float64(2) {
		t.Fatalf("expected y=2, got %#v", j["y"])
	}

	if err := j.Scan(nil); err != nil {
		t.Fatalf("scan nil failed: %v", err)
	}
	if j != nil {
		t.Fatalf("expected nil map after nil scan, got %#v", j)
	}

	if err := j.Scan(42); err == nil {
		t.Fatal("expected unsupported type error")
	}
}
