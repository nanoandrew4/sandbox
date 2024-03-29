package datastructs

import "errors"

type sNode[T any] struct {
	val  T
	next *sNode[T]
}

type Stack[T any] struct {
	head *sNode[T]
}

func (s *Stack[T]) Push(val T) {
	if s.head == nil {
		s.head = &sNode[T]{val: val}
		return
	}
	s.head = &sNode[T]{val: val, next: s.head}
}

func (s *Stack[T]) Pop() (rVal T, err error) {
	if s.head == nil {
		return rVal, errors.New("stack is empty")
	}
	rVal = s.head.val
	s.head = s.head.next
	return rVal, nil
}

func (s *Stack[T]) Peek() (rVal T, err error) {
	if s.head == nil {
		return rVal, errors.New("stack is empty")
	}
	return s.head.val, nil
}
