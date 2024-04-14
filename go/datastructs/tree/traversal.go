package tree

import "sandbox/types"

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
