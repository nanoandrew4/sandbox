package tree

import (
	"sandbox/types"
)

type BinaryTree[T types.Sortable, R binaryNode[T]] interface {
	Insert(valuesToInsert ...T)

	root() R
	setRoot(newRoot R)
}

type baseBinaryTree[T types.Sortable, R binaryNode[T]] struct {
	rootNode R
}

func (tree *baseBinaryTree[T, R]) root() R {
	return tree.rootNode
}

func (tree *baseBinaryTree[T, R]) setRoot(newRoot R) {
	tree.rootNode = newRoot
}

func (tree *baseBinaryTree[T, R]) Equals(bt2 BinaryTree[T, R]) bool {
	if bt2 == nil {
		return false
	}
	if !areNodesEqual[T](tree.root(), bt2.root()) {
		return false
	}
	return true
}

func areNodesEqual[T types.Sortable](n1, n2 binaryNode[T]) bool {
	if isNodeNil(n1) && isNodeNil(n2) {
		return true
	} else if isNodeNil(n1) || isNodeNil(n2) {
		return false
	} else if n1.val() != n2.val() {
		return false
	}
	return areNodesEqual(n1.left(), n2.left()) && areNodesEqual(n1.right(), n2.right())
}
