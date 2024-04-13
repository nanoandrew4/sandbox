package tree

import (
	"sandbox/datastructs"
	"sandbox/types"
)

type TreeType int

const (
	preOrderTree = TreeType(iota)
	inOrderTree
	postOrderTree
)

type NodeTraversalFunc[T types.Sortable] func(node binaryNode[T]) (continueTraversal bool)

type BinaryTree[T types.Sortable] struct {
	tType TreeType
	root  *binaryTreeNode[T]
}

func (bt *BinaryTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if bt.root == nil {
		bt.root = &binaryTreeNode[T]{nodeVal: valuesToInsert[0]}
		valuesToInsert = valuesToInsert[1:]
	}

	if len(valuesToInsert) > 0 {
		var valIdx int
		q := &datastructs.Queue[binaryNode[T]]{}
		bt.walkBreadthFirstAndRunFunc(q, bt.root, func(node binaryNode[T]) bool {
			if isNodeNil(node.left()) {
				node.setLeft(&binaryTreeNode[T]{nodeVal: valuesToInsert[valIdx]})
				q.Enqueue(node)
				valIdx++
				return valIdx != len(valuesToInsert)
			} else if isNodeNil(node.right()) {
				node.setRight(&binaryTreeNode[T]{nodeVal: valuesToInsert[valIdx]})
				q.Enqueue(node)
				valIdx++
				return valIdx != len(valuesToInsert)
			}
			return true
		})
	}
}

func (bt *BinaryTree[T]) WalkDepthFirst(f NodeTraversalFunc[T]) {
	bt.walkDepthFirstAndRunFunc(bt.root, f)
}

func (bt *BinaryTree[T]) walkDepthFirstAndRunFunc(node binaryNode[T], f NodeTraversalFunc[T]) {
	if isNodeNil(node) {
		return
	}

	switch bt.tType {
	case preOrderTree:
		if !f(node) {
			return
		}
		bt.walkDepthFirstAndRunFunc(node.left(), f)
		bt.walkDepthFirstAndRunFunc(node.right(), f)
	case inOrderTree:
		bt.walkDepthFirstAndRunFunc(node.left(), f)
		if !f(node) {
			return
		}
		bt.walkDepthFirstAndRunFunc(node.right(), f)
	case postOrderTree:
		bt.walkDepthFirstAndRunFunc(node.left(), f)
		bt.walkDepthFirstAndRunFunc(node.right(), f)
		if !f(node) {
			return
		}
	}
}

func (bt *BinaryTree[T]) walkBreadthFirstAndRunFunc(q *datastructs.Queue[binaryNode[T]], node binaryNode[T], f NodeTraversalFunc[T]) {
	if node == nil {
		return
	}

	q.Enqueue(node.left())
	q.Enqueue(node.right())
	if !f(node) {
		return
	}
	nextNode, dequeueErr := q.Dequeue() // if dequeue fails, nextNode should be nil, which will be caught with the first if statement
	for nextNode == nil && dequeueErr == nil {
		nextNode, dequeueErr = q.Dequeue()
	}
	bt.walkBreadthFirstAndRunFunc(q, nextNode, f)
}

func (bt *BinaryTree[T]) Contains(val T) bool {
	var found bool
	bt.walkDepthFirstAndRunFunc(bt.root, func(node binaryNode[T]) (continueTraversal bool) {
		if !isNodeNil(node) && node.val() == val {
			found = true
			return false
		}
		return true
	})
	return found
}

func (bt *BinaryTree[T]) Equals(bt2 *BinaryTree[T]) bool {
	if bt2 == nil {
		return false
	}
	if !areNodesEqual[T](bt.root, bt2.root) {
		return false
	}
	return true
}

func areNodesEqual[T types.Sortable](n1, n2 binaryNode[T]) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	} else if n1.val() != n2.val() {
		return false
	}
	return areNodesEqual(n1.left(), n2.left()) && areNodesEqual(n1.right(), n2.right())
}

func (bt *BinaryTree[T]) Delete(val T) bool {
	return bt.deleteVal(nil, bt.root, val)
}

func (bt *BinaryTree[T]) deleteVal(parent, node binaryNode[T], val T) bool {
	if isNodeNil(node) {
		return false
	}

	if node.val() == val {
		if parent == nil {
			bt.root = nil
		} else if isNodeNil(node.left()) && isNodeNil(node.right()) {
			if parent.left() == node {
				parent.setLeft(nil)
			} else {
				parent.setRight(nil)
			}
		} else if !isNodeNil(node.left()) {
			bubbleUpLeftSideValues(parent, node)
		} else if !isNodeNil(node.right()) {
			bubbleUpRightSideValues(parent, node)
		}
		return true
	}
	return bt.deleteVal(node, node.left(), val) || bt.deleteVal(node, node.right(), val)
}

func bubbleUpLeftSideValues[T types.Sortable](parent, node binaryNode[T]) *T {
	if node == nil {
		return nil
	}

	originalNodeVal := node.val()
	lChildVal := bubbleUpLeftSideValues(node, node.left())
	if lChildVal != nil {
		node.setVal(*lChildVal)
	} else {
		parent.setLeft(nil) // delete last node on branch
	}

	return &originalNodeVal
}

func bubbleUpRightSideValues[T types.Sortable](parent, node binaryNode[T]) *T {
	if isNodeNil(node) {
		return nil
	}

	originalNodeVal := node.val()
	rChildVal := bubbleUpRightSideValues(node, node.right())
	if rChildVal != nil {
		node.setVal(*rChildVal)
	} else {
		parent.setRight(nil) // delete last node on branch
	}

	return &originalNodeVal
}
