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

type UnorderedBinaryTree[T types.Sortable] struct {
	baseBinaryTree[T, *binaryTreeNode[T]]
	tType TreeType
}

func (bt *UnorderedBinaryTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if bt.root() == nil {
		bt.setRoot(&binaryTreeNode[T]{nodeVal: valuesToInsert[0]})
		valuesToInsert = valuesToInsert[1:]
	}

	if len(valuesToInsert) > 0 {
		var valIdx int
		q := &datastructs.Queue[binaryNode[T]]{}
		bt.walkBreadthFirstAndRunFunc(q, bt.root(), func(node binaryNode[T]) bool {
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

func (bt *UnorderedBinaryTree[T]) WalkDepthFirst(f NodeTraversalFunc[T]) {
	bt.walkDepthFirstAndRunFunc(bt.root(), f)
}

func (bt *UnorderedBinaryTree[T]) walkDepthFirstAndRunFunc(node binaryNode[T], f NodeTraversalFunc[T]) {
	if isNodeNil(node) {
		return
	}

	switch bt.tType {
	case preOrderTree:
		TraversePreOrder(node, f)
	case inOrderTree:
		TraverseInOrder(node, f)
	case postOrderTree:
		TraversePostOrder(node, f)
	}
}

func (bt *UnorderedBinaryTree[T]) walkBreadthFirstAndRunFunc(q *datastructs.Queue[binaryNode[T]], node binaryNode[T], f NodeTraversalFunc[T]) {
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

func (bt *UnorderedBinaryTree[T]) Contains(val T) bool {
	var found bool
	bt.walkDepthFirstAndRunFunc(bt.root(), func(node binaryNode[T]) (continueTraversal bool) {
		if !isNodeNil(node) && node.val() == val {
			found = true
			return false
		}
		return true
	})
	return found
}

func (bt *UnorderedBinaryTree[T]) Delete(val T) bool {
	return bt.deleteVal(nil, bt.root(), val)
}

func (bt *UnorderedBinaryTree[T]) deleteVal(parent, node binaryNode[T], val T) bool {
	if isNodeNil(node) {
		return false
	}

	if node.val() == val {
		if isNodeNil(parent) {
			bt.setRoot(nil)
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
	if isNodeNil(node) {
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
