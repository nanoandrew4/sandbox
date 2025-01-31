package tree

import (
	"fmt"
	"sandbox/types"
)

type baseSortedBinaryTree[T types.Sortable, R binaryNode[T]] struct {
	baseBinaryTree[T, R]
}

func (tree *baseSortedBinaryTree[T, R]) ToOrderedArray() []T {
	var arr []T
	TraverseInOrder[T](tree.root(), func(node binaryNode[T]) bool {
		arr = append(arr, node.val())
		return true
	})
	return arr
}

func (tree *baseSortedBinaryTree[T, R]) Find(val T) binaryNode[T] {
	_, node := findParentAndNodeByVal(nil, tree.root(), val)
	return node
}

func (tree *baseSortedBinaryTree[T, R]) Contains(val T) bool {
	return tree.Find(val) != nil
}

func findParentAndNodeByVal[T types.Sortable](parent, node binaryNode[T], val T) (p, n binaryNode[T]) {
	if isNodeNil[T](node) {
		return nil, nil
	} else if node.val() == val {
		return parent, node
	} else if val < node.val() {
		return findParentAndNodeByVal(node, node.left(), val)
	} else {
		return findParentAndNodeByVal(node, node.right(), val)
	}
}

func deleteBinaryNode[T types.Sortable, R binaryNode[T]](tree BinaryTree[T, R], val T) (deleted bool, nodeToReBalanceFrom R) {
	nodeToDeleteParent, nodeToDelete := findParentAndNodeByVal(nil, tree.root(), val)

	if isNodeNil(nodeToDelete) {
		return
	}

	if isNodeNil(nodeToDeleteParent) { // height of tree decreases
		tree.setRoot(nodeToReBalanceFrom) // nodeToReBalanceFrom will be nil at this point if R is a pointer type
		return true, nodeToReBalanceFrom
	}

	var dirInParent direction
	if nodeToDeleteParent.right() == nodeToDelete {
		dirInParent = right
	}

	if isNodeNil(nodeToDelete.left()) && isNodeNil(nodeToDelete.right()) { // height of tree decreases
		nodeToDeleteParent.setChildDir(nil, dirInParent)
		fmt.Println("deleting leaf")
		return true, nodeToDeleteParent.(R)
	} else if nodeToDeleteParent.childDir(dirInParent) == nodeToDelete && isNodeNil(nodeToDelete.childDir(1-dirInParent)) { // height of subtree decreases
		fmt.Printf("replace child %v with same dir nephew %v\n", nodeToDelete.val(), nodeToDelete.childDir(dirInParent).val())
		nodeToDeleteParent.setChildDir(nodeToDelete.childDir(dirInParent), dirInParent)
		return true, nodeToDeleteParent.(R)
	} else {
		dirInParent = right
		if !isNodeNil(nodeToDelete.left()) { // prefer descending left side, if possible
			dirInParent = left
		}

		var l, r T
		if !isNodeNil(nodeToDelete.left()) {
			l = nodeToDelete.left().val()
		}
		if !isNodeNil(nodeToDelete.right()) {
			r = nodeToDelete.right().val()
		}
		fmt.Printf("will delete node %v with l-child %v and r-child %v\n", nodeToDelete.val(), l, r)

		parentOfNodeToSwap, nodeToSwap := findFarthestNodeInDir(nodeToDelete, nodeToDelete.childDir(dirInParent), 1-dirInParent)
		nodeToSwapChildDir := nodeToSwap.childDir(dirInParent)
		parentOfNodeToSwap.setChildDir(nodeToSwapChildDir, 1-dirInParent)
		nodeToDelete.setVal(nodeToSwap.val())
		if nodeToSwapChildDirWithParent, castOk := nodeToSwapChildDir.(binaryNodeWithParent[T]); castOk {
			nodeToSwapChildDirWithParent.setParent(parentOfNodeToSwap.(binaryNodeWithParent[T]))
		}
		return true, nodeToSwap.(R)
	}
}

func findFarthestNodeInDir[T types.Sortable](parent, node binaryNode[T], dir direction) (p, n binaryNode[T]) {
	if isNodeNil(node) {
		return nil, nil
	}

	p, n = findFarthestNodeInDir(node, node.childDir(dir), dir)
	if isNodeNil(p) && isNodeNil(n) {
		return parent, node
	}
	return p, n
}
