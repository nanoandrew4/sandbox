package tree

import "sandbox/datastructs"

type TreeType int

const (
	preOrderTree = TreeType(iota)
	inOrderTree
	postOrderTree
)

type NodeTraversalFunc[T comparable] func(node *btNode[T]) (continueTraversal bool)

type btNode[T comparable] struct {
	left, right *btNode[T]
	val         T
}

type BinaryTree[T comparable] struct {
	tType TreeType
	root  *btNode[T]
}

func (bt *BinaryTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if bt.root == nil {
		bt.root = &btNode[T]{val: valuesToInsert[0]}
		valuesToInsert = valuesToInsert[1:]
	}

	if len(valuesToInsert) > 0 {
		var valIdx int
		q := &datastructs.Queue[*btNode[T]]{}
		bt.walkBreadthFirstAndRunFunc(q, bt.root, func(node *btNode[T]) bool {
			if node.left == nil {
				node.left = &btNode[T]{val: valuesToInsert[valIdx]}
				q.Enqueue(node)
				valIdx++
				return valIdx != len(valuesToInsert)
			} else if node.right == nil {
				node.right = &btNode[T]{val: valuesToInsert[valIdx]}
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

func (bt *BinaryTree[T]) walkDepthFirstAndRunFunc(node *btNode[T], f NodeTraversalFunc[T]) {
	if node == nil {
		return
	}

	switch bt.tType {
	case preOrderTree:
		if !f(node) {
			return
		}
		bt.walkDepthFirstAndRunFunc(node.left, f)
		bt.walkDepthFirstAndRunFunc(node.right, f)
	case inOrderTree:
		bt.walkDepthFirstAndRunFunc(node.left, f)
		if !f(node) {
			return
		}
		bt.walkDepthFirstAndRunFunc(node.right, f)
	case postOrderTree:
		bt.walkDepthFirstAndRunFunc(node.left, f)
		bt.walkDepthFirstAndRunFunc(node.right, f)
		if !f(node) {
			return
		}
	}
}

func (bt *BinaryTree[T]) walkBreadthFirstAndRunFunc(q *datastructs.Queue[*btNode[T]], node *btNode[T], f NodeTraversalFunc[T]) {
	if node == nil {
		return
	}

	q.Enqueue(node.left)
	q.Enqueue(node.right)
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
	bt.walkDepthFirstAndRunFunc(bt.root, func(node *btNode[T]) (continueTraversal bool) {
		if node != nil && node.val == val {
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
	if !areNodesEqual(bt.root, bt2.root) {
		return false
	}
	return true
}

func areNodesEqual[T comparable](n1, n2 *btNode[T]) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	} else if n1.val != n2.val {
		return false
	}
	return areNodesEqual(n1.left, n2.left) && areNodesEqual(n1.right, n2.right)
}

func (bt *BinaryTree[T]) Delete(val T) bool {
	return bt.deleteVal(nil, bt.root, val)
}

func (bt *BinaryTree[T]) deleteVal(parent, node *btNode[T], val T) bool {
	if node == nil {
		return false
	}

	if node.val == val {
		if parent == nil {
			bt.root = nil
		} else if node.left == nil && node.right == nil {
			if parent.left == node {
				parent.left = nil
			} else {
				parent.right = nil
			}
		} else if node.left != nil {
			bubbleUpLeftSideValues(parent, node)
		} else if node.right != nil {
			bubbleUpRightSideValues(parent, node)
		}
		return true
	}
	return bt.deleteVal(node, node.left, val) || bt.deleteVal(node, node.right, val)
}

func bubbleUpLeftSideValues[T comparable](parent, node *btNode[T]) *T {
	if node == nil {
		return nil
	}

	originalNodeVal := node.val
	lChildVal := bubbleUpLeftSideValues(node, node.left)
	if lChildVal != nil {
		node.val = *lChildVal
	} else {
		parent.left = nil // delete last node on branch
	}

	return &originalNodeVal
}

func bubbleUpRightSideValues[T comparable](parent, node *btNode[T]) *T {
	if node == nil {
		return nil
	}

	originalNodeVal := node.val
	lChildVal := bubbleUpRightSideValues(node, node.right)
	if lChildVal != nil {
		node.val = *lChildVal
	} else {
		parent.right = nil // delete last node on branch
	}

	return &originalNodeVal
}
