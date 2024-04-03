package algorithms

import "sandbox/types"

func BinarySearch[T types.Sortable](arr []T, valToSearch T) bool {
	midIdx := len(arr) / 2
	if len(arr) > 1 {
		if arr[midIdx] > valToSearch {
			return BinarySearch(arr[:midIdx], valToSearch)
		} else if arr[midIdx] < valToSearch {
			return BinarySearch(arr[midIdx:], valToSearch)
		}
	}
	return arr[midIdx] == valToSearch
}

func NonRecursiveBinarySearch[T types.Sortable](arr []T, valToSearch T) bool {
	var lowIdx, midIdx int
	highIdx := len(arr)

	var midVal T
	for lowIdx < highIdx {
		midIdx = (highIdx + lowIdx) / 2
		midVal = arr[midIdx]
		if midVal == valToSearch {
			return true
		} else if valToSearch > midVal {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx
		}
	}
	return false
}
