package algorithms

import (
	"fmt"
	"math/rand"
	"testing"
)

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

func BenchmarkBinarySearch(b *testing.B) {
	for _, numbersToSort := range []int{1e6, 1e7, 1e8} {
		testValues := makeRange(0, numbersToSort)
		b.Run(fmt.Sprintf("Recursive/%d", numbersToSort), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = BinarySearch(testValues, rand.Int()%numbersToSort)
			}
		})
		b.Run(fmt.Sprintf("NonRecursive/%d", numbersToSort), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NonRecursiveBinarySearch(testValues, rand.Int()%numbersToSort)
			}
		})
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
