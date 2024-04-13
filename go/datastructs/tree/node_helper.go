package tree

import "sandbox/types"

func isNodeNil[T types.Sortable](node binaryNode[T]) bool {
	switch node.(type) {
	case *binaryTreeNode[T]:
		return node.(*binaryTreeNode[T]) == nil
	case *binaryTreeNodeWithParent[T]:
		return node.(*binaryTreeNodeWithParent[T]) == nil
	case *avlTNode[T]:
		return node.(*avlTNode[T]) == nil
	case *rbTNode[T]:
		return node.(*rbTNode[T]) == nil
	}
	return true
}
