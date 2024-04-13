package tree

import "sandbox/types"

type binaryNode[T types.Sortable] interface {
	left() binaryNode[T]
	setLeft(newLeft binaryNode[T])
	right() binaryNode[T]
	setRight(newRight binaryNode[T])
	val() T
	setVal(val T)
	childDir(dir direction) binaryNode[T]
	setChildDir(newChild binaryNode[T], dir direction)
}

type binaryTreeNode[T types.Sortable] struct {
	children [2]binaryNode[T]
	nodeVal  T
}

func (node *binaryTreeNode[T]) left() binaryNode[T] {
	return node.children[left]
}

func (node *binaryTreeNode[T]) setLeft(newLeft binaryNode[T]) {
	node.children[left] = newLeft
}

func (node *binaryTreeNode[T]) right() binaryNode[T] {
	return node.children[right]
}

func (node *binaryTreeNode[T]) setRight(newRight binaryNode[T]) {
	node.children[right] = newRight
}

func (node *binaryTreeNode[T]) val() T {
	return node.nodeVal
}

func (node *binaryTreeNode[T]) setVal(val T) {
	node.nodeVal = val
}

func (node *binaryTreeNode[T]) childDir(dir direction) binaryNode[T] {
	return node.children[dir]
}

func (node *binaryTreeNode[T]) setChildDir(newChild binaryNode[T], dir direction) {
	node.children[dir] = newChild
}
