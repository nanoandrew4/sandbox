package algorithms

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		testArr := generateRandomNumbers(1000)
		QuickSort(testArr)
		for j := 0; j < len(testArr)-1; j++ {
			if testArr[j] > testArr[j+1] {
				t.Error(fmt.Sprintf("sorting error at index %d", j))
				break
			}
		}
	}
}

func TestBalancedQuickSort(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		testArr := generateRandomNumbers(1000)
		BalancedQuickSort(testArr)
		for j := 0; j < len(testArr)-1; j++ {
			if testArr[j] > testArr[j+1] {
				t.Error(fmt.Sprintf("sorting error at index %d", j))
				break
			}
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	numOfArraysToGenerateMap := map[int]int{
		1e4: 10000,
		1e6: 1000,
		1e7: 100,
	}
	for _, numbersToSort := range []int{1e4, 1e6, 1e7} {
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
				QuickSort(testValues)
			}
		})
		b.Run(fmt.Sprintf("Balanced/%d", numbersToSort), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				testValues := make([]int, len(aggregatedTestValues[0]))
				copy(testValues, aggregatedTestValues[i%len(aggregatedTestValues)])
				b.StartTimer()
				BalancedQuickSort(testValues)
			}
		})
	}
}
