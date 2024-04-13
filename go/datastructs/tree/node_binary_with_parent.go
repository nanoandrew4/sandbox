package tree

import "sandbox/types"

type binaryNodeWithParent[T types.Sortable] interface {
	binaryNode[T]
	parent() binaryNodeWithParent[T]
	setParent(binaryNodeWithParent[T])
}

type binaryTreeNodeWithParent[T types.Sortable] struct {
	binaryNode[T]
	nodeParent binaryNodeWithParent[T]
}

func (node *binaryTreeNodeWithParent[T]) parent() binaryNodeWithParent[T] {
	return node.nodeParent
}

func (node *binaryTreeNodeWithParent[T]) setParent(parent binaryNodeWithParent[T]) {
	node.nodeParent = parent
}

func swapChild[T types.Sortable](parent, oldChild, newChild binaryNodeWithParent[T]) {
	if parent.left() == oldChild {
		parent.setLeft(newChild)
	} else {
		parent.setRight(newChild)
	}
}
