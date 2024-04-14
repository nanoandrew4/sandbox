package tree

import "sandbox/types"

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
