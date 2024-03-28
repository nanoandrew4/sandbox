package datastructs

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr [5]int      // size defined at compile time, immutable
	fmt.Println(arr[0]) // can print, array has 5 zero values initialized

	// Cannot append to array, if uncommented will not compile
	//arr = append(arr, 1)
	//newArr := append(arr, 1)
}
