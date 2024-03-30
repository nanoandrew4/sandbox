package datastructs

import "errors"

type RingBuffer[T any] struct {
	arr              []T
	headIdx, tailIdx int
	length           int
}

func (rb *RingBuffer[T]) Enqueue(val T) {
	if cap(rb.arr) <= rb.length {
		rb.resize(rb.length + 1)
	}
	rb.arr[rb.tailIdx] = val
	rb.length++
	rb.tailIdx++
	if rb.tailIdx >= len(rb.arr) {
		rb.tailIdx %= len(rb.arr)
	}
}

func (rb *RingBuffer[T]) Dequeue() (retVal T, err error) {
	if rb.length == 0 {
		return retVal, errors.New("no element to dequeue")
	}
	retVal = rb.arr[rb.headIdx]
	rb.length--
	rb.headIdx++
	if rb.headIdx >= len(rb.arr) {
		rb.headIdx %= len(rb.arr)
	}
	return retVal, nil
}

func (rb *RingBuffer[T]) Push(val T) {
	if cap(rb.arr) <= rb.length {
		rb.resize(rb.length + 1)
	}
	rb.length++
	rb.headIdx--
	if rb.headIdx < 0 {
		rb.headIdx += len(rb.arr)
	}
	rb.arr[rb.headIdx] = val
}

func (rb *RingBuffer[T]) Pop() (retVal T, err error) {
	if rb.length == 0 {
		return retVal, errors.New("no element to pop")
	}
	rb.length--
	rb.tailIdx--
	if rb.tailIdx < 0 {
		rb.tailIdx += len(rb.arr)
	}
	retVal = rb.arr[rb.tailIdx]
	return retVal, nil
}

func (rb *RingBuffer[T]) Peek() (retVal T, err error) {
	if rb.length == 0 {
		return retVal, errors.New("no element to peek")
	}

	return rb.arr[rb.tailIdx-1], nil
}

func (rb *RingBuffer[T]) resize(minCap int) {
	newArr := make([]T, max(cap(rb.arr)*2, minCap))
	copy(newArr, append(rb.arr[rb.headIdx:], rb.arr[:rb.tailIdx]...))
	rb.headIdx = 0
	rb.tailIdx = rb.length
	rb.arr = newArr
}
