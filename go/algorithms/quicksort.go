package algorithms

import "sandbox/types"

func QuickSort[T types.Sortable](arr []T) {
	qs(arr, 0, len(arr)-1)
}

func qs[T types.Sortable](arr []T, lo, hi int) {
	if lo >= hi {
		return
	}
	piv := partition(arr, lo, hi)
	qs(arr, lo, piv-1)
	qs(arr, piv+1, hi)
}

func partition[T types.Sortable](arr []T, lo, hi int) (pivIdx int) {
	idx := lo - 1
	pivot := arr[hi]
	var tmp T
	for ; lo < hi; lo++ {
		if arr[lo] <= pivot {
			idx++
			if lo != idx {
				tmp = arr[lo]
				arr[lo] = arr[idx]
				arr[idx] = tmp
			}
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot
	return idx
}

// For some reason this was the quicksort implementation I came up with when trying to implement it from memory, but
// is less performant than the standard implementation
func BalancedQuickSort[T types.Sortable](arr []T) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[len(arr)/2]
	var lIdx, rIdx int
	var tmp T
	for rIdx = len(arr) - 1; lIdx < rIdx; {
		if arr[lIdx] >= pivot && arr[rIdx] <= pivot {
			tmp = arr[lIdx]
			arr[lIdx] = arr[rIdx]
			arr[rIdx] = tmp
			lIdx++
		} else if arr[lIdx] < pivot {
			lIdx++
		} else if arr[rIdx] > pivot {
			rIdx--
		}
	}
	BalancedQuickSort(arr[:lIdx])
	BalancedQuickSort(arr[lIdx:])
}
