package tree

import (
	"sandbox/types"
)

type BinarySearchTree[T types.Sortable] struct {
	root *binaryTreeNode[T]
}

func (bst *BinarySearchTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if bst.root == nil {
		bst.root = &binaryTreeNode[T]{nodeVal: valuesToInsert[0]}
		valuesToInsert = valuesToInsert[1:]
	}

	for _, val := range valuesToInsert {
		insertValue(bst.root, val)
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

func (bst *BinarySearchTree[T]) Contains(val T) bool {
	_, node := findParentAndNodeByVal(nil, bst.root, val)
	return node != nil
}

func findParentAndNodeByVal[T types.Sortable](parent, node binaryNode[T], val T) (p, n binaryNode[T]) {
	if isNodeNil(node) {
		return nil, nil
	} else if node.val() == val {
		return parent, node
	} else if val < node.val() {
		return findParentAndNodeByVal(node, node.left(), val)
	} else {
		return findParentAndNodeByVal(node, node.right(), val)
	}
}

func (bst *BinarySearchTree[T]) Delete(val T) bool {
	nodeToDeleteParent, nodeToDelete := findParentAndNodeByVal(nil, bst.root, val)

	if nodeToDelete == nil {
		return false
	} else if nodeToDeleteParent == nil {
		bst.root = nil
	} else if nodeToDelete.left() == nil && nodeToDelete.right() == nil {
		if nodeToDeleteParent.left() == nodeToDelete {
			nodeToDeleteParent.setLeft(nil)
		} else {
			nodeToDeleteParent.setRight(nil)
		}
	} else if nodeToDelete.right() != nil {
		_, minNode := findMinNode(nodeToDelete, nodeToDelete.right())
		nodeToDelete.setVal(minNode.val())
		minNode.setVal(minNode.right().val())
		minNode.setLeft(minNode.right().left())
		minNode.setRight(minNode.right().right())
	} else {
		nodeToDelete.setVal(nodeToDelete.left().val())
		nodeToDelete.setLeft(nodeToDelete.left().left())
		nodeToDelete.setRight(nodeToDelete.left().right())
	}

	return true
}

func findMinNode[T types.Sortable](parent, node binaryNode[T]) (p, n binaryNode[T]) {
	if node == nil {
		return nil, nil
	}

	p, n = findMinNode(node, node.left())
	if p == nil && n == nil {
		return parent, node
	}
	return p, n
}

func (bst *BinarySearchTree[T]) ToOrderedArray() []T {
	var arr []T
	walkInOrder(bst.root, func(node binaryNode[T]) {
		arr = append(arr, node.val())
	})
	return arr
}

func walkInOrder[T types.Sortable](node binaryNode[T], f func(node binaryNode[T])) {
	if node == nil {
		return
	}
	walkInOrder(node.left(), f)
	f(node)
	walkInOrder(node.right(), f)
}
