package datastructs

import "testing"

func TestRingBuffer_Enqueue(t *testing.T) {
	rb := RingBuffer[int]{}
	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Enqueue(3)
	rb.Enqueue(4)

	if val, err := rb.Peek(); val != 4 || err != nil {
		t.Error("expected peeked value to be 4")
	}
}

func TestRingBuffer_Push(t *testing.T) {
	rb := RingBuffer[int]{}
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)
	rb.Push(4)

	if val, err := rb.Peek(); val != 1 || err != nil {
		t.Error("expected peeked value to be 1")
	}
}
