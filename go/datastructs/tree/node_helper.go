package tree

import "sandbox/types"

func isNodeNil[T types.Sortable](node binaryNode[T]) bool {
	switch node.(type) {
	case *binaryTreeNode[T]:
		return node.(*binaryTreeNode[T]) == nil
	case *binaryTreeNodeWithParent[T]:
		return node.(*binaryTreeNodeWithParent[T]) == nil
	case binaryTreeNodeWithParent[T]:
		return node.(binaryTreeNodeWithParent[T]) == binaryTreeNodeWithParent[T]{}
	case *avlTNode[T]:
		return node.(*avlTNode[T]) == nil
	case avlTNode[T]:
		return node.(avlTNode[T]) == avlTNode[T]{}
	case *rbTNode[T]:
		return node.(*rbTNode[T]) == nil
	case rbTNode[T]:
		return node.(rbTNode[T]) == rbTNode[T]{}
	}
	return true
}

func castAndReturnNode[T types.Sortable, R binaryNode[T]](node binaryNode[T]) R {
	if isNodeNil[T](node) {
		var defaultVal R
		return defaultVal
	}
	return node.(R)
}
