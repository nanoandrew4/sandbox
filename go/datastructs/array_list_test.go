package datastructs

import "testing"

func TestArrayList_Prepend(t *testing.T) {
	al := &ArrayList[int]{}

	al.Prepend(1)
	al.Prepend(2, 3)
	al.Prepend(4)

	if val, err := al.Get(0); val != 4 || err != nil {
		t.Error("expected first element of the array list to be 4")
	}
}

func TestArrayList_Append(t *testing.T) {
	al := &ArrayList[int]{}

	al.Append(1)
	al.Append(2, 3)
	al.Append(4)

	if val, err := al.Get(0); val != 1 || err != nil {
		t.Error("expected first element of the array list to be 1")
	}
}

func TestArrayList_Pop(t *testing.T) {
	al := &ArrayList[int]{}

	al.Append(1)
	al.Append(2, 3)
	al.Append(4)

	if val, err := al.Pop(); val != 4 || err != nil {
		t.Error("expected popped element to be 1")
	}
	if val, err := al.Pop(); val != 3 || err != nil {
		t.Error("expected popped element to be 2")
	}
}

func TestArrayList_Peek(t *testing.T) {
	al := &ArrayList[int]{}

	al.Append(1)
	al.Append(2, 3)
	al.Append(4)

	if val, err := al.Peek(); val != 4 || err != nil {
		t.Error("expected popped element to be 1")
	}

	if val, err := al.Peek(); val != 4 || err != nil {
		t.Error("expected popped element to be 1 again")
	}
}
