package tree

import (
	"sandbox/types"
)

type bstNode[T types.Sortable] struct {
	left, right *bstNode[T]
	val         T
}

type BinarySearchTree[T types.Sortable] struct {
	root *bstNode[T]
}

func (bst *BinarySearchTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if bst.root == nil {
		bst.root = &bstNode[T]{val: valuesToInsert[0]}
		valuesToInsert = valuesToInsert[1:]
	}

	for _, val := range valuesToInsert {
		insertValue(bst.root, val)
	}
}

func insertValue[T types.Sortable](node *bstNode[T], val T) {
	if val < node.val {
		if node.left != nil {
			insertValue(node.left, val)
		} else {
			node.left = &bstNode[T]{val: val}
		}
	} else {
		if node.right != nil {
			insertValue(node.right, val)
		} else {
			node.right = &bstNode[T]{val: val}
		}
	}
}

func (bst *BinarySearchTree[T]) Contains(val T) bool {
	_, node := findParentAndNodeByVal(nil, bst.root, val)
	return node != nil
}

func findParentAndNodeByVal[T types.Sortable](parent, node *bstNode[T], val T) (p, n *bstNode[T]) {
	if node == nil {
		return nil, nil
	} else if node.val == val {
		return parent, node
	} else if val < node.val {
		return findParentAndNodeByVal(node, node.left, val)
	} else {
		return findParentAndNodeByVal(node, node.right, val)
	}
}

func (bst *BinarySearchTree[T]) Delete(val T) bool {
	nodeToDeleteParent, nodeToDelete := findParentAndNodeByVal(nil, bst.root, val)

	if nodeToDelete == nil {
		return false
	} else if nodeToDeleteParent == nil {
		bst.root = nil
	} else if nodeToDelete.left == nil && nodeToDelete.right == nil {
		if nodeToDeleteParent.left == nodeToDelete {
			nodeToDeleteParent.left = nil
		} else {
			nodeToDeleteParent.right = nil
		}
	} else if nodeToDelete.right != nil {
		_, minNode := findMinNode(nodeToDelete, nodeToDelete.right)
		nodeToDelete.val = minNode.val
		minNode.val = minNode.right.val
		minNode.left = minNode.right.left
		minNode.right = minNode.right.right
	} else {
		nodeToDelete.val = nodeToDelete.left.val
		nodeToDelete.left = nodeToDelete.left.left
		nodeToDelete.right = nodeToDelete.left.right
	}

	return true
}

func findMinNode[T types.Sortable](parent, node *bstNode[T]) (p, n *bstNode[T]) {
	if node == nil {
		return nil, nil
	}

	p, n = findMinNode(node, node.left)
	if p == nil && n == nil {
		return parent, node
	}
	return p, n
}

func (bst *BinarySearchTree[T]) ToOrderedArray() []T {
	var arr []T
	walkInOrder(bst.root, func(node *bstNode[T]) {
		arr = append(arr, node.val)
	})
	return arr
}

func walkInOrder[T types.Sortable](node *bstNode[T], f func(node *bstNode[T])) {
	if node == nil {
		return
	}
	walkInOrder(node.left, f)
	f(node)
	walkInOrder(node.right, f)
}
