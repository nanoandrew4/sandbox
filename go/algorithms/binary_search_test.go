package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	if !BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7) {
		t.Error("expected to find 7 in ordered list of numbers")
	}
	if BinarySearch([]int{1, 2, 3, 4, 5, 6, 8, 9, 10}, 7) {
		t.Error("expected not to find 7 in ordered list of numbers")
	}
}

func TestNonRecursiveBinarySearch(t *testing.T) {
	if !NonRecursiveBinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7) {
		t.Error("expected to find 7 in ordered list of numbers")
	}
	if NonRecursiveBinarySearch([]int{1, 2, 3, 4, 5, 6, 8, 9, 10}, 7) {
		t.Error("expected not to find 7 in ordered list of numbers")
	}
}
