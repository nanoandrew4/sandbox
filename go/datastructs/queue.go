package datastructs

import "errors"

type qNode[T comparable] struct {
	val  T
	next *qNode[T]
}

type Queue[T comparable] struct {
	head, tail *qNode[T]
}

func (q *Queue[T]) Enqueue(val T) {
	if q.tail == nil {
		q.head = &qNode[T]{val: val}
		q.tail = q.head
		return
	}
	q.tail.next = &qNode[T]{val: val}
	q.tail = q.tail.next
}

func (q *Queue[T]) Dequeue() (rVal T, err error) {
	if q.head == nil {
		return rVal, errors.New("queue is empty")
	}
	rVal = q.head.val
	q.head = q.head.next
	return rVal, nil
}

func (q *Queue[T]) Peek() (rVal T, err error) {
	if q.head == nil {
		return rVal, errors.New("queue is empty")
	}
	return q.head.val, nil
}
