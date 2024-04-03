package datastructs

import (
	"slices"
	"testing"
)

func TestSinglyLinkedList_Prepend(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(1)
	ll.Prepend(3)
	ll.Prepend(5)
	if !slices.Equal(ll.ToArray(), []int{5, 3, 1}) {
		t.Fatal("expected linked list to be 5, 3, 1")
	}
}

func TestSinglyLinkedList_Append(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Append(1)
	ll.Append(3)
	ll.Append(5)
	if !slices.Equal(ll.ToArray(), []int{1, 3, 5}) {
		t.Fatal("expected linked list to be 1, 3, 5")
	}
}
func TestSinglyLinkedList_Contains(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Append(1)
	ll.Append(3)
	ll.Append(5)
	if !ll.Contains(3) {
		t.Fatal("expected linked list to contain 3")
	}
	if ll.Contains(2) {
		t.Fatal("expected linked list to not contain 2")
	}
}

func TestSinglyLinkedList_Size(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Append(1)
	ll.Append(3)
	ll.Append(5)
	if ll.Size() != 3 {
		t.Fatal("expected linked list size to be 3")
	}

	ll.Reset()
	if ll.Size() != 0 {
		t.Fatal("expected linked list size to be 0")
	}
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	ll := SinglyLinkedList[string]{}
	ll.Prepend("one")
	ll.Prepend("two")
	ll.Prepend("three")
	ll.Reverse()
	if !slices.Equal(ll.ToArray(), []string{"one", "two", "three"}) {
		t.Fatal("expected linked list to be ordered as one, two, three")
	}

	ll.Reverse()
	ll.Reverse()
	if !slices.Equal(ll.ToArray(), []string{"one", "two", "three"}) {
		t.Fatal("expected linked list to be ordered as one, two, three")
	}
}

func TestSinglyLinkedList_Get(t *testing.T) {
	ll := SinglyLinkedList[string]{}
	ll.Prepend("one")
	ll.Prepend("two")
	ll.Prepend("three")
	val, err := ll.Get(0)
	if err != nil || val != "three" {
		t.Fatal("expected retrieved value to be \"one\"")
	}
	val, err = ll.Get(-1)
	if err == nil || val != "" {
		t.Fatal("expected error and default value for searching with negative index")
	}

	val, err = ll.Get(10)
	if err == nil || val != "" {
		t.Fatal("expected error and default value for searching out of bounds")
	}
}

func TestSinglyLinkedList_Delete(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	if ll.Delete(0) != nil {
		t.Fatal("error deleting first element of list")
	}
	if !slices.Equal(ll.ToArray(), []int{2, 3, 4, 5}) {
		t.Fatal("expected 1 to be deleted from list")
	}

	if ll.Delete(3) != nil {
		t.Fatal("error deleting last element of list")
	}
	if !slices.Equal(ll.ToArray(), []int{2, 3, 4}) {
		t.Fatal("expected 5 to be deleted from list")
	}

	if ll.Delete(1) != nil {
		t.Fatal("error deleting middle element of list")
	}
	if !slices.Equal(ll.ToArray(), []int{2, 4}) {
		t.Fatal("expected 3 to be deleted from list")
	}

	if ll.Delete(-1) == nil {
		t.Fatal("expected error deleting at negative index")
	}
}
