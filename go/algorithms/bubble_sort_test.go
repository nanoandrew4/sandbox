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
				t.Fatal(fmt.Sprintf("sorting error at index %d", j-1))
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
				t.Fatal(fmt.Sprintf("sorting error at index %d", j-1))
			}
		}
	}
}

func TestSlowBubbleSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		testValues := generateRandomNumbers(1000)
		CleanFastBubbleSort(testValues)
		for j := 0; j < len(testValues)-1; j++ {
			if testValues[j] > testValues[j+1] {
				t.Fatal(fmt.Sprintf("sorting error at index %d", j-1))
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	numOfArraysToGenerateMap := map[int]int{
		1e2: 100000,
		1e3: 10000,
		1e4: 1000,
	}
	for _, numbersToSort := range []int{1e2, 1e3, 1e4} {
		var aggregatedTestValues [][]int
		for t := 0; t < numOfArraysToGenerateMap[numbersToSort]; t++ {
			testValues := generateRandomNumbers(numbersToSort)
			aggregatedTestValues = append(aggregatedTestValues, testValues)
		}

		b.Run(fmt.Sprintf("Std/%d", numbersToSort), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				testValues := make([]int, len(aggregatedTestValues[0]))
				copy(testValues, aggregatedTestValues[i%len(aggregatedTestValues)])
				b.StartTimer()
				BubbleSort(testValues)
			}
		})
		b.Run(fmt.Sprintf("Fast/%d", numbersToSort), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				testValues := make([]int, len(aggregatedTestValues[0]))
				copy(testValues, aggregatedTestValues[i%len(aggregatedTestValues)])
				b.StartTimer()
				FastBubbleSort(testValues)
			}
		})
		b.Run(fmt.Sprintf("CleanFastBubbleSort/%d", numbersToSort), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				testValues := make([]int, len(aggregatedTestValues[0]))
				copy(testValues, aggregatedTestValues[i%len(aggregatedTestValues)])
				b.StartTimer()
				CleanFastBubbleSort(testValues)
			}
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
