package datastructs

import "errors"

type sllNode[T comparable] struct {
	val  T
	next *sllNode[T]
}

type SinglyLinkedList[T comparable] struct {
	head *sllNode[T]
}

func (ll *SinglyLinkedList[T]) Prepend(val T) {
	if ll.head == nil {
		ll.head = &sllNode[T]{val: val}
		return
	}
	ll.head = &sllNode[T]{val: val, next: ll.head}
}

func (ll *SinglyLinkedList[T]) Append(val T) {
	if ll.head == nil {
		ll.head = &sllNode[T]{val: val}
		return
	}
	node := ll.head
	for ; node.next != nil; node = node.next {
	}

	node.next = &sllNode[T]{val: val}
}

func (ll *SinglyLinkedList[T]) Get(idx int) (T, error) {
	for i, node := 0, ll.head; node != nil; i, node = i+1, node.next {
		if i == idx && node != nil {
			return node.val, nil
		}
	}
	var rVal T
	return rVal, errors.New("out of bounds")
}

func (ll *SinglyLinkedList[T]) Delete(idx int) error {
	if idx == 0 {
		if ll.head != nil {
			ll.head = ll.head.next
		}
		return nil
	}

	var prevNode *sllNode[T]
	for i, node := 0, ll.head; node != nil; i, prevNode, node = i+1, node, node.next {
		if i == idx && node != nil && prevNode != nil {
			prevNode.next = node.next
			node.next = nil
			return nil
		}
	}
	return errors.New("out of bounds")
}

func (ll *SinglyLinkedList[T]) Contains(val T) bool {
	for node := ll.head; node != nil; node = node.next {
		if node.val == val {
			return true
		}
	}
	return false
}

func (ll *SinglyLinkedList[T]) Size() (size int) {
	for node := ll.head; node != nil; node, size = node.next, size+1 {
	}
	return
}

func (ll *SinglyLinkedList[T]) Reset() {
	ll.head = nil
}

func (ll *SinglyLinkedList[T]) Reverse() {
	var prevNode, tmpNode *sllNode[T]
	for node := ll.head; node != nil; prevNode, node = node, tmpNode {
		tmpNode = node.next
		node.next = prevNode
	}
	ll.head = prevNode
}

func (ll *SinglyLinkedList[T]) ToArray() (outArr []T) {
	for node := ll.head; node != nil; node = node.next {
		outArr = append(outArr, node.val)
	}
	return
}
