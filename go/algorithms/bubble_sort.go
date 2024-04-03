package algorithms

import "sandbox/types"

func BubbleSort[T types.Sortable](arr []T) {
	var tmp T
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i; j++ {
			if j+1 < len(arr) && arr[j+1] < arr[j] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

func FastBubbleSort[T types.Sortable](arr []T) {
	var tmp T
	var arrLen = len(arr)
	for j := 0; j < arrLen; j++ {
		if j+1 < arrLen && arr[j+1] < arr[j] {
			tmp = arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = tmp
		}
	}
	for i := arrLen - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if j+1 < i && arr[j+1] < arr[j] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

func CleanFastBubbleSort[T types.Sortable](arr []T) {
	var tmp T
	var arrLen = len(arr)
	innerSort := func(upperBound int) {
		for j := 0; j < upperBound; j++ {
			if j+1 < upperBound && arr[j+1] < arr[j] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	innerSort(arrLen)
	for i := arrLen - 1; i > 0; i-- {
		innerSort(i)
	}
}
