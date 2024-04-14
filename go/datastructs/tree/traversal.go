package tree

import (
	"sandbox/datastructs"
	"sandbox/types"
)

type NodeTraversalFunc[T types.Sortable] func(node binaryNode[T]) (continueTraversal bool)

func TraversePreOrder[T types.Sortable](node binaryNode[T], f NodeTraversalFunc[T]) {
	if isNodeNil(node) {
		return
	} else if !f(node) {
		return
	}
	TraversePreOrder(node.left(), f)
	TraversePreOrder(node.right(), f)
}

func TraverseInOrder[T types.Sortable](node binaryNode[T], f NodeTraversalFunc[T]) {
	if isNodeNil(node) {
		return
	}

	TraverseInOrder(node.left(), f)
	if !f(node) {
		return
	}
	TraverseInOrder(node.right(), f)
}

func TraversePostOrder[T types.Sortable](node binaryNode[T], f NodeTraversalFunc[T]) {
	if isNodeNil(node) {
		return
	}

	TraversePostOrder(node.left(), f)
	TraversePostOrder(node.right(), f)
	if !f(node) {
		return
	}
}

func TraverseBreadthFirst[T types.Sortable](node binaryNode[T], f NodeTraversalFunc[T]) {
	q := &datastructs.Queue[binaryNode[T]]{}
	walkBreadthFirstAndRunFunc(q, node, f)
}

func walkBreadthFirstAndRunFunc[T types.Sortable](q *datastructs.Queue[binaryNode[T]], node binaryNode[T], f NodeTraversalFunc[T]) {
	if isNodeNil(node) {
		return
	}

	q.Enqueue(node.left())
	q.Enqueue(node.right())
	if !f(node) {
		return
	}
	nextNode, dequeueErr := q.Dequeue() // if dequeue fails, nextNode should be nil, which will be caught with the first if statement
	for isNodeNil(nextNode) && dequeueErr == nil {
		nextNode, dequeueErr = q.Dequeue()
	}
	walkBreadthFirstAndRunFunc(q, nextNode, f)
}
