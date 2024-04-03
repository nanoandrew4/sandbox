package datastructs

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	var slice []int

	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fatal("expected print to panic")
			}
		}()
		fmt.Println(slice[0])
	}()

	slice = append(slice, 1, 2, 3, 4, 5) // can append to slice, since it is not a fixed size
	if slice[0] != 1 {
		t.Fatal("expected first element to equal 1")
	}

	newSlice := slice
	newSlice[0] = -1
	if slice[0] != -1 {
		t.Fatal("expected first element in original array to be mutated")
	}

	newSlice = append(newSlice, 6)
	newSlice[0] = 1
	if len(slice)+1 != len(newSlice) {
		t.Fatal("newSlice should have 1 more element")
	}

	for i := 0; i < 5; i++ {
		sPtr := &slice[i]
		nsPtr := &newSlice[i]
		if sPtr != nsPtr {
			t.Fatal("both slices should point to the same memory addresses for the first 5 elements")
		}
	}

	slice = append(slice, 7)
	for i := 0; i < 6; i++ { // the original slice modifies the copy, removing the 6 we appended before
		sPtr := &slice[i]
		nsPtr := &newSlice[i]
		if sPtr != nsPtr {
			t.Fatal("both slices should point to the same memory addresses for the first 6 elements")
		}
	}

	newSlice = newSlice[:3]
	newSlice = append(newSlice, -4, -5)
	for i := 0; i < 5; i++ { // the copied slice modifies the original, causing its fourth and fifth elements to also be modified
		sPtr := &slice[i]
		nsPtr := &newSlice[i]
		if sPtr != nsPtr {
			t.Fatal("both slices should point to the same memory addresses for the first 5 elements")
		}
	}

	func(s []int) {
		newSlice = append(newSlice, 6, 7)
	}(newSlice)
	if len(slice) == len(newSlice) {
		t.Fatal("appending to a slice inside a function should modify the slice outside of the function")
	}

	newSlice = append(newSlice, []int{8}...)
	for i := 0; i < 5; i++ {
		sPtr := &slice[i]
		nsPtr := &newSlice[i]
		if sPtr == nsPtr {
			t.Fatal("both slices should not point to the same memory addresses for the first 5 elements, since appending two slices results in a new underlying array")
		}
	}
}
