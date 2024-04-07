package list

import "errors"

type dllNode[T comparable] struct {
	val        T
	prev, next *dllNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head *dllNode[T]
}

func (ll *DoublyLinkedList[T]) Prepend(val T) {
	if ll.head == nil {
		ll.head = &dllNode[T]{val: val}
		return
	}
	ll.head = &dllNode[T]{val: val, next: ll.head}
	ll.head.next.prev = ll.head
}

func (ll *DoublyLinkedList[T]) Append(val T) {
	if ll.head == nil {
		ll.head = &dllNode[T]{val: val}
		return
	}
	node := ll.head
	for ; node.next != nil; node = node.next {
	}

	node.next = &dllNode[T]{val: val, prev: node}
}

func (ll *DoublyLinkedList[T]) Get(idx int) (T, error) {
	for i, node := 0, ll.head; node != nil; i, node = i+1, node.next {
		if i == idx {
			return node.val, nil
		}
	}
	var rVal T
	return rVal, errors.New("out of bounds")
}

func (ll *DoublyLinkedList[T]) Delete(idx int) error {
	if idx == 0 {
		if ll.head != nil {
			ll.head = ll.head.next
		}
		return nil
	}

	var prevNode *dllNode[T]
	for i, node := 0, ll.head; node != nil; i, prevNode, node = i+1, node, node.next {
		if i == idx && prevNode != nil {
			prevNode.next = node.next
			if node.next != nil {
				node.next.prev = prevNode
			}
			node.next = nil
			node.prev = nil
			return nil
		}
	}
	return errors.New("out of bounds")
}

func (ll *DoublyLinkedList[T]) Contains(val T) bool {
	for node := ll.head; node != nil; node = node.next {
		if node.val == val {
			return true
		}
	}
	return false
}

func (ll *DoublyLinkedList[T]) Size() (size int) {
	for node := ll.head; node != nil; node, size = node.next, size+1 {
	}
	return
}

func (ll *DoublyLinkedList[T]) Reset() {
	ll.head = nil
}

func (ll *DoublyLinkedList[T]) Reverse() {
	var prevNode, tmpNode *dllNode[T]
	for node := ll.head; node != nil; prevNode, node = node, tmpNode {
		tmpNode = node.next
		node.next = prevNode
		node.prev = tmpNode
	}
	ll.head = prevNode
}

func (ll *DoublyLinkedList[T]) ToArray() (outArr []T) {
	for node := ll.head; node != nil; node = node.next {
		outArr = append(outArr, node.val)
	}
	return
}

func (ll *DoublyLinkedList[T]) WalkOutAndBack(f func(nodeVal T)) {
	var node *dllNode[T]
	for node = ll.head; node != nil; {
		f(node.val)
		if node.next != nil {
			node = node.next
		} else {
			break
		}
	}
	for ; node != nil; node = node.prev {
		f(node.val)
	}
}
