package datastructs

import "errors"

type sNode[T comparable] struct {
	val  T
	next *sNode[T]
}

type Stack[T comparable] struct {
	head *sNode[T]
}

func (q *Stack[T]) Push(val T) {
	if q.head == nil {
		q.head = &sNode[T]{val: val}
		return
	}
	q.head = &sNode[T]{val: val, next: q.head}
}

func (q *Stack[T]) Pop() (rVal T, err error) {
	if q.head == nil {
		return rVal, errors.New("queue is empty")
	}
	rVal = q.head.val
	q.head = q.head.next
	return rVal, nil
}

func (q *Stack[T]) Peek() (rVal T, err error) {
	if q.head == nil {
		return rVal, errors.New("stack is empty")
	}
	return q.head.val, nil
}
