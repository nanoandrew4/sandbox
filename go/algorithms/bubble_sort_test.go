package algorithms

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testValues := generateRandomNumbers(1000)
		BubbleSort(testValues)
		for j := 0; j < len(testValues)-1; j++ {
			if testValues[j] > testValues[j+1] {
				t.Error(fmt.Sprintf("sorting error at index %d", j-1))
				break
			}
		}
	}
}

func TestFastBubbleSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testValues := generateRandomNumbers(1000)
		FastBubbleSort(testValues)
		for j := 0; j < len(testValues)-1; j++ {
			if testValues[j] > testValues[j+1] {
				t.Error(fmt.Sprintf("sorting error at index %d", j-1))
				break
			}
		}
	}
}

func TestSlowBubbleSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testValues := generateRandomNumbers(1000)
		SlowBubbleSort(testValues)
		for j := 0; j < len(testValues)-1; j++ {
			if testValues[j] > testValues[j+1] {
				t.Error(fmt.Sprintf("sorting error at index %d", j-1))
				break
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, numbersToSort := range []int{1000, 10000, 100000} {
		b.Run(fmt.Sprintf("%d", numbersToSort), func(b *testing.B) {
			testValues := generateRandomNumbers(numbersToSort)
			b.ResetTimer()
			BubbleSort(testValues)
		})
	}
}

func BenchmarkFastBubbleSort(b *testing.B) {
	for _, numbersToSort := range []int{1000, 10000, 100000} {
		b.Run(fmt.Sprintf("%d", numbersToSort), func(b *testing.B) {
			testValues := generateRandomNumbers(numbersToSort)
			b.ResetTimer()
			FastBubbleSort(testValues)
		})
	}
}

func BenchmarkSlowBubbleSort(b *testing.B) {
	for _, numbersToSort := range []int{1000, 10000, 100000} {
		b.Run(fmt.Sprintf("%d", numbersToSort), func(b *testing.B) {
			testValues := generateRandomNumbers(numbersToSort)
			b.ResetTimer()
			SlowBubbleSort(testValues)
		})
	}
}

func generateRandomNumbers(n int) []int {
	randomNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		randomNumbers[i] = rand.Int() % n
	}
	return randomNumbers
}
