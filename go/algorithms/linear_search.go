package algorithms

func LinearFind[T comparable](arr []T, valToFind T) bool {
	return IndexOf(arr, valToFind) != -1
}

func IndexOf[T comparable](arr []T, valToFind T) int {
	for idx, val := range arr {
		if val == valToFind {
			return idx
		}
	}
	return -1
}
