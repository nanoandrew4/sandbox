package datastructs

import "testing"

func TestStack(t *testing.T) {
	q := Stack[int]{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if val, err := q.Peek(); val != 3 || err != nil {
		t.Fatal("expected first peeked value to be 3")
	}
	if val, err := q.Pop(); val != 3 || err != nil {
		t.Fatal("expected first pop value to be 3")
	}
	if val, err := q.Peek(); val != 2 || err != nil {
		t.Fatal("expected second peeked value to be 2")
	}
	if val, err := q.Pop(); val != 2 || err != nil {
		t.Fatal("expected second pop value to be 2")
	}
	if val, err := q.Peek(); val != 1 || err != nil {
		t.Fatal("expected third peeked value to be 1")
	}
	if val, err := q.Pop(); val != 1 || err != nil {
		t.Fatal("expected third pop value to be 1")
	}
	if val, err := q.Peek(); val != 0 || err == nil {
		t.Fatal("expected fourth peeked value to return error, since queue should be empty")
	}
	if val, err := q.Pop(); val != 0 || err == nil {
		t.Fatal("expected fourth pop value to return error, since queue should be empty")
	}
}
