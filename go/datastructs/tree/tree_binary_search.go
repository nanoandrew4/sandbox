package tree

import (
	"sandbox/types"
)

type BinarySearchTree[T types.Sortable] struct {
	baseSortedBinaryTree[T, *binaryTreeNode[T]]
}

func (bst *BinarySearchTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if bst.root() == nil {
		bst.setRoot(&binaryTreeNode[T]{nodeVal: valuesToInsert[0]})
		valuesToInsert = valuesToInsert[1:]
	}

	for _, val := range valuesToInsert {
		insertValue(bst.root(), val)
	}
}

func insertValue[T types.Sortable](node binaryNode[T], val T) {
	if val < node.val() {
		if node.left() != nil {
			insertValue(node.left(), val)
		} else {
			node.setLeft(&binaryTreeNode[T]{nodeVal: val})
		}
	} else {
		if node.right() != nil {
			insertValue(node.right(), val)
		} else {
			node.setRight(&binaryTreeNode[T]{nodeVal: val})
		}
	}
}

func (bst *BinarySearchTree[T]) Delete(val T) bool {
	deleted, _ := deleteBinaryNode[T, *binaryTreeNode[T]](bst, val)
	return deleted
}
