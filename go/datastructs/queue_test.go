package datastructs

import "testing"

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if val, err := q.Peek(); val != 1 || err != nil {
		t.Fatal("expected first peeked value to be 1")
	}
	if val, err := q.Dequeue(); val != 1 || err != nil {
		t.Fatal("expected first dequeued value to be 1")
	}
	if val, err := q.Peek(); val != 2 || err != nil {
		t.Fatal("expected second peeked value to be 2")
	}
	if val, err := q.Dequeue(); val != 2 || err != nil {
		t.Fatal("expected second dequeued value to be 2")
	}
	if val, err := q.Peek(); val != 3 || err != nil {
		t.Fatal("expected third peeked value to be 3")
	}
	if val, err := q.Dequeue(); val != 3 || err != nil {
		t.Fatal("expected third dequeued value to be 3")
	}
	if val, err := q.Peek(); val != 0 || err == nil {
		t.Fatal("expected fourth peeked value to return error, since queue should be empty")
	}
	if val, err := q.Dequeue(); val != 0 || err == nil {
		t.Fatal("expected fourth dequeued value to return error, since queue should be empty")
	}
}
